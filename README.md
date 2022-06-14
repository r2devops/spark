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
         "/home/thomas/workspace/r2devops/jobs/Dockerfile"
      ],
      "Data": null
   },
   "golang": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/configuration/configuration.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/gitlab.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/gitlab_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/localCrawler.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/localCrawler_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/ports/file_service.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/ports/version_analyzer.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/request.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/request_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/analyzer.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/analyzer_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/models/messages.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/models/version.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/models/version_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/parser.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/parser_test.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/versions.go",
         "/home/thomas/workspace/r2devops/jobs/crawler/versions/versions_test.go",
         "/home/thomas/workspace/r2devops/jobs/db/db.go",
         "/home/thomas/workspace/r2devops/jobs/db/labels.go",
         "/home/thomas/workspace/r2devops/jobs/db/licenseTypes.go",
         "/home/thomas/workspace/r2devops/jobs/db/logger.go",
         "/home/thomas/workspace/r2devops/jobs/go.mod",
         "/home/thomas/workspace/r2devops/jobs/handlers/context/Sort.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/context/context.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/handler.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/handler_test.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/job.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/job_test.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/pipeline.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/ports/crawler.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/project.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/request.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/response.go",
         "/home/thomas/workspace/r2devops/jobs/handlers/routes.go",
         "/home/thomas/workspace/r2devops/jobs/job/helper.go",
         "/home/thomas/workspace/r2devops/jobs/job/job.go",
         "/home/thomas/workspace/r2devops/jobs/job/update.go",
         "/home/thomas/workspace/r2devops/jobs/licenses/licenses.go",
         "/home/thomas/workspace/r2devops/jobs/licenses/projects.go",
         "/home/thomas/workspace/r2devops/jobs/main.go",
         "/home/thomas/workspace/r2devops/jobs/middlewares/jobQuery.go",
         "/home/thomas/workspace/r2devops/jobs/middlewares/kratosAuth.go",
         "/home/thomas/workspace/r2devops/jobs/middlewares/kratosAuth_test.go",
         "/home/thomas/workspace/r2devops/jobs/models/models.go",
         "/home/thomas/workspace/r2devops/jobs/pipelines/pipelines.go",
         "/home/thomas/workspace/r2devops/jobs/router/router.go",
         "/home/thomas/workspace/r2devops/jobs/router/validator.go",
         "/home/thomas/workspace/r2devops/jobs/s3/s3.go",
         "/home/thomas/workspace/r2devops/jobs/s3/s3_test.go",
         "/home/thomas/workspace/r2devops/jobs/utils/config.go",
         "/home/thomas/workspace/r2devops/jobs/utils/errors.go",
         "/home/thomas/workspace/r2devops/jobs/utils/gitlab_helper.go",
         "/home/thomas/workspace/r2devops/jobs/utils/gitlab_helper_test.go",
         "/home/thomas/workspace/r2devops/jobs/utils/job.go",
         "/home/thomas/workspace/r2devops/jobs/utils/job_test.go",
         "/home/thomas/workspace/r2devops/jobs/utils/kratos/kratos.go",
         "/home/thomas/workspace/r2devops/jobs/utils/pagination.go",
         "/home/thomas/workspace/r2devops/jobs/utils/pipeline.go",
         "/home/thomas/workspace/r2devops/jobs/utils/token/token.go"
      ],
      "Data": [
         {
            "Name": "version",
            "Value": "1.18",
            "FilePath": "/home/thomas/workspace/r2devops/jobs/go.mod"
         },
         {
            "Name": "module_name",
            "Value": "gitlab.com/r2devops/jobs",
            "FilePath": "/home/thomas/workspace/r2devops/jobs/go.mod"
         }
      ]
   },
   "hcl": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/terraform/production/.terraform.lock.hcl",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/backend.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/database.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/jobs.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/locals.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/provider.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/s3.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/variables.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/.terraform.lock.hcl",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/backend.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/database.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/jobs.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/locals.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/provider.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/s3.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/variables.tf"
      ],
      "Data": null
   },
   "javascript": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasDiscriminator.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasDocumentSchema.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasExample.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasOpFormDataConsumeCheck.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasOpIdUnique.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasOpParams.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasOpSecurityDefined.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasOpSuccessResponse.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasPathParam.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasSchema.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasTagDefined.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/oasUnusedComponent.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/refSiblings.js",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/custom-functions/typedEnum.js"
      ],
      "Data": null
   },
   "json": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/.spectral.json",
         "/home/thomas/workspace/r2devops/jobs/.stoplight/styleguide.json",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/.terraform/terraform.tfstate",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/.terraform/terraform.tfstate",
         "/home/thomas/workspace/r2devops/jobs/test/testdata/job_test_data_job_list.json",
         "/home/thomas/workspace/r2devops/jobs/test/testdata/job_test_data_job_list_filtered_not_deprecated.json",
         "/home/thomas/workspace/r2devops/jobs/test/testdata/job_test_data_job_list_filtered_owner.json",
         "/home/thomas/workspace/r2devops/jobs/test/testdata/job_test_data_job_list_filtered_owner_search_not_deprecated.json",
         "/home/thomas/workspace/r2devops/jobs/test/testdata/job_test_data_job_list_filtered_search.json"
      ],
      "Data": null
   },
   "markdown": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/.gitlab/issue_templates/Bug.md",
         "/home/thomas/workspace/r2devops/jobs/.gitlab/issue_templates/Default.md",
         "/home/thomas/workspace/r2devops/jobs/.gitlab/issue_templates/Feature.md",
         "/home/thomas/workspace/r2devops/jobs/.gitlab/merge_request_templates/Default.md",
         "/home/thomas/workspace/r2devops/jobs/README.md",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/.terraform/providers/registry.terraform.io/scaleway/scaleway/2.2.1/linux_amd64/CHANGELOG.md",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/.terraform/providers/registry.terraform.io/scaleway/scaleway/2.2.1/linux_amd64/README.md",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/.terraform/providers/registry.terraform.io/scaleway/scaleway/2.2.1/linux_amd64/CHANGELOG.md",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/.terraform/providers/registry.terraform.io/scaleway/scaleway/2.2.1/linux_amd64/README.md"
      ],
      "Data": null
   },
   "perl": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/terraform/production/backend.tf",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/backend.tf"
      ],
      "Data": null
   },
   "yaml": {
      "FilesMatch": [
         "/home/thomas/workspace/r2devops/jobs/.gitlab-ci.yml",
         "/home/thomas/workspace/r2devops/jobs/docker-compose.yml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/Chart.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/configmap.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/cronjob.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/deployment.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/hook_jobs.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/ingress.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/secrets.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/service.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/templates/tests/test-connection.yaml",
         "/home/thomas/workspace/r2devops/jobs/helm-chart/jobs/values.yaml",
         "/home/thomas/workspace/r2devops/jobs/reference/jobs.v1.yaml",
         "/home/thomas/workspace/r2devops/jobs/terraform/production/conf/jobs-values.yaml",
         "/home/thomas/workspace/r2devops/jobs/terraform/staging/conf/jobs-values.yaml"
      ],
      "Data": null
   }
}" action=main context=spark_cli
```
