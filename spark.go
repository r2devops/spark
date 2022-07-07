package spark

import (
	"path"
	"regexp"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var logger = logrus.WithField("context", "spark")

type Configuration struct {
	Languages []struct {
		Name      string   `yaml:"name"`
		Filenames []string `yaml:"filenames"`
		FileRegex []*regexp.Regexp
		Extract   []struct {
			Name          string `yaml:"name"`
			File          string `yaml:"file"`
			Regex         string `yaml:"regex"`
			CompiledRegex *regexp.Regexp
		} `yaml:"extract"`
	} `yaml:"languages"`
}

type LanguageData struct {
	FilesMatch []string
	Data       map[string]*ExtractedData
}

type ExtractedData struct {
	Value    string
	FilePath string
}

func Analyze(filePath string, content []byte, conf *Configuration) map[string]*LanguageData {

	// Initialize logger
	l := logger.WithFields(logrus.Fields{
		"action": "file_analyze",
	})

	result := make(map[string]*LanguageData, 0)

	// For each languages
	for _, language := range conf.Languages {

		languageData := LanguageData{}

		// Get file base nanme (without full path)
		fileBaseName := path.Base(filePath)

		// Check if one filename regex match
		for _, filename := range language.Filenames {

			// Check if filename match
			match, err := regexp.MatchString(filename, fileBaseName)
			if err != nil {
				l.WithField("filename_regex", filename).WithError(err).Error("unable to parse regex, skip it")
				continue
			}
			if match {
				l.WithField("filename_regex", filename).Debug("regex match !")
				// Add matching file to
				languageData.FilesMatch = append(languageData.FilesMatch, filePath)
				break
			}
		}

		// Even if there is no match in filenames, it may be a match in data
		// extraction file
		languageData.Data = make(map[string]*ExtractedData, 0)
		for _, extract := range language.Extract {

			// If extraction file is the current file and complied regex exists
			if fileBaseName == extract.File && extract.CompiledRegex != nil {

				result := extract.CompiledRegex.FindSubmatch(content)
				if result != nil && len(result) > 1 {
					extractedData := ExtractedData{
						Value:    string(result[1]),
						FilePath: filePath,
					}
					languageData.Data[extract.Name] = &extractedData
				} else {
					l.WithFields(logrus.Fields{
						"file":  extract.File,
						"regex": extract.Regex,
					}).Debug("rexeg doesn't match")

				}
			}
		}

		// If no filematch neither data => next language
		if len(languageData.Data) == 0 && languageData.FilesMatch == nil {
			continue
		}

		// Append to existing key or create new one
		if _, ok := result[language.Name]; ok {
			result[language.Name].FilesMatch = append(result[language.Name].FilesMatch, languageData.FilesMatch...)
			for k, v := range languageData.Data {
				result[language.Name].Data[k] = v
			}
		} else {
			result[language.Name] = &languageData
		}
	}

	l.WithField("result", result).Debug("language detection done")

	return result
}

func ParseConf(fileContent []byte) (*Configuration, error) {

	// Initialize logger
	l := logger.WithFields(logrus.Fields{
		"action": "parse_conf",
	})

	var conf *Configuration

	err := yaml.Unmarshal(fileContent, &conf)
	if err != nil {
		l.WithError(err).Error("unable to unmarshal yaml configuration file")
		return conf, err
	}

	// Prepare compilation of all regex to extract data
	for iLanguage, language := range conf.Languages {
		for iExtract, extract := range language.Extract {
			regex, err := regexp.Compile(extract.Regex)
			if err != nil {
				l.WithField("regex", extract.Regex).WithError(err).Error("unable to data extraction compile regex, skip it")
				continue
			}
			conf.Languages[iLanguage].Extract[iExtract].CompiledRegex = regex
			l.WithField("*regex", regex).Debug("regex is compiled !")
		}
	}
	return conf, nil
}
