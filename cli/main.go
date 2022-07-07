package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/r2devops/spark"
	"github.com/sirupsen/logrus"
)

const configPath = "conf.yaml"

var logger = logrus.WithField("context", "spark_cli")

// Search recursivelly for all files in a directory and return all paths
func getFileList(dirPath string, subDirsToSkip []string) ([]string, error) {

	l := logger.WithFields(logrus.Fields{
		"action": "get_file_list",
	})

	var filesFound []string
	sort.Strings(subDirsToSkip)

	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			l.WithField("path", path).WithError(err).Error("failure accessing the path")
			return err
		}

		if info.IsDir() {
			index := sort.SearchStrings(subDirsToSkip, info.Name())
			if index != len(subDirsToSkip) && subDirsToSkip[index] == info.Name() {
				l.WithField("directory", info.Name()).Debug("skipping dir without error")
				return filepath.SkipDir
			}
		} else if info.Mode().IsRegular() {
			l.Debug("adding file to list: ", path)
			filesFound = append(filesFound, path)
		}

		return nil
	})

	if err != nil {
		l.WithField("path", dirPath).WithError(err).Error("error walking the path")
		return []string{}, err
	}

	return filesFound, err
}

func main() {

	// Initialize logger
	l := logger.WithFields(logrus.Fields{
		"action": "main",
	})

	dirsToSkip := []string{".git", "node_modules", "vendor"}

	// Get file path as arg
	if len(os.Args) < 2 {
		l.Error("path to directory to analyze must be specified as arg")
		os.Exit(1)
	}
	dirPath := os.Args[1]

	// Open configuration file and get content
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		l.WithError(err).Error("unable to open file")
		os.Exit(1)
	}
	fileContent, err := ioutil.ReadFile(configPath)
	conf, err := spark.ParseConf(fileContent)
	if err != nil {
		os.Exit(1)
	}
	l.WithField("conf", conf).Debug("configuration content")
	l.Info("configuration initialized")

	// Get a list of all files from the root directory
	files, err := getFileList(dirPath, dirsToSkip)
	if err != nil {
		l.WithError(err).Error("error when trying to list recursivelly all files")
		os.Exit(1)
	}

	// Analyze all techs
	techs := map[string]*spark.LanguageData{}
	for _, filePath := range files {

		// Get the file content
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			l.WithField("file", filePath).WithError(err).Warning("unable to open file")
			continue
		}

		// Analyze file
		fileTechs := spark.Analyze(filePath, fileContent, conf)
		l.WithField("file", filePath).Debug("analyze of file done")

		// Append it to result
		for tech, content := range fileTechs {

			if _, ok := techs[tech]; ok {
				techs[tech].FilesMatch = append(techs[tech].FilesMatch, content.FilesMatch...)
				for k, v := range content.Data {
					techs[tech].Data[k] = v
				}
			} else {
				techs[tech] = content
			}
		}
	}

	// Display result
	result, _ := json.MarshalIndent(techs, "", "   ")
	l.Info("Analysis result: " + string(result))
}
