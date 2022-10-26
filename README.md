# Spark

This project detect languages/technologies used in file(s) and is extract data.

## Run it

```shell
go run cli/main.go <path to file or directory to analyze>
```

## Configuration

Configuration describing how to detect languages and what data extract is
written in the file `conf.yaml` following this format:

```yaml
# languages:
#   - name: golang                  ==> language name
#     filenames:
#       - '^go.mod$"                ==> regex in filename
#       - '^.*\.go$'
#     extract:
#       - name: version             ==> name of paramter to extract
#         file: 'go.mod"            ==> exact filename (not regex) of file containing the paramter to extract
#         regex: '^go ([0-9\.]+)'   ==> regex to extract the parameter with exactly one submatch "()'
```

## Example

```shell
go run cli/main.go ~/workspace/r2devops/jobs

time="2022-06-21T11:31:09+02:00" level=info msg="Analysis result: {
   "dockerfile": {
      "FilesMatch": [
         "REDACTED"
      ],
      "Data": null
   },
   "golang": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED",
         "REDACTED"
      ],
      "Data": [
         {
            "Name": "version",
            "Value": "1.18",
            "FilePath": "REDACTED"
         },
         {
            "Name": "module_name",
            "Value": "gitlab.com/r2devops/jobs",
            "FilePath": "REDACTED"
         }
      ]
   },
   "hcl": {
      "FilesMatch": [
         "REDACTED"
      ],
      "Data": null
   },
   "javascript": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED"
      ],
      "Data": null
   },
   "json": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED"
      ],
      "Data": null
   },
   "markdown": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED"
      ],
      "Data": null
   },
   "perl": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED"
      ],
      "Data": null
   },
   "yaml": {
      "FilesMatch": [
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED",
         "REDACTED"
      ],
      "Data": null
   }
}" action=main context=spark_cli
```
