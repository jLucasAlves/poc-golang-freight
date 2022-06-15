# Template-golang

[![CircleCI](https://circleci.com/gh/dafiti-group/golang-template-project/tree/main.svg?style=svg)](https://circleci.com/gh/dafiti-group/golang-template-project/tree/main) [![Go Reference](https://pkg.go.dev/badge/go.dev/doc/.svg)](https://go.dev/doc/)
## Sample application for Go lang framework

## How to execute template local

To run the project locally, replace the variables with the tag `${{values.component_id}}` with the name of the desired Application and `${{values.description}}` with description application.

## Links devops Tools

| Tool         | URL                                                                                                                            |
| :----------- | :----------------------------------------------------------------------------------------------------------------------------- |
| CicleCI      | <https://app.circleci.com/pipelines/github/dafiti-group/template-golang>                                                       |
| SonarCube    | <http://sonarqube.live.br.dafiti.io/dashboard?id=template-golang>                                                              |
| Argo QA      | <https://argocd.dafiti.la/applications/template-golang-br-qa>                                                                  |
| Argo Live    | <https://argocd.dafiti.la/applications/template-golang-br-live>                                                                |
| GrayLog QA   | <https://graylog.qa.dafiti.local/search?q=kubernetes_namespace_name%3A%22template-golang%22&rangetype=relative&relative=300>   |
| GrayLog Live | <https://graylog.live.dafiti.local/search?q=kubernetes_namespace_name%3A%22template-golang%22&rangetype=relative&relative=300> |
| Instana      | <https://apm-dafiti.instana.io/#/services>                                                                                     |

## Hosts

| Env        | Path             | URL                                                                    |
| :--------- | :--------------- | :--------------------------------------------------------------------- |
| Local      | root             | <http://localhost:8080>                                                |
| Local      | Liveness Health  | <http://localhost:8080/health-check/liveness>                          |
| Local      | Readiness Health | <http://localhost:8080/health-check/readiness>                         |
|            |                  |                                                                        |
| QA         | root             | <https://template-golang.eks.qa.dafiti.local>                          |
| QA         | Liveness Health  | <https://template-golang.eks.qa.dafiti.local/health-check/liveness>    |
| QA         | Readiness Health | <https://template-golang.eks.qa.dafiti.local/health-check/readiness>   |
|            |                  |                                                                        |
| Production | root             | <https://template-golang.eks.live.dafiti.local>                        |
| Production | Liveness Health  | <https://template-golang.eks.live.dafiti.local/health-check/liveness>  |
| Production | Readiness Health | <https://template-golang.eks.live.dafiti.local/health-check/readiness> |

## Below is the config you need to do before execute Project

- Be sure about your permissions in [THIS CONFLUENSE PAGE](https://dafiti.jira.com/wiki/spaces/DFTEC/pages/3247013947/Desenvolvimento+local+DOCKER-DAFITI)

## Development (outside dev-container VScode)

The requirement to use this repo is `docker` and `docker-compose >=v1.27`, if you need
to install this plugin, follow these steps:

- install [Docker](https://docs.docker.com/engine/install/ubuntu/)
- install [Docker compose](https://docs.docker.com/compose/install/)

and all you need to run is...

```sh
docker-compose up app
```

If your application is not up and running, try to double check the envs in `docker-compose.yaml` file

## Testing

To test your application you need to run `docker-compose run --rm test`
and this command will use the correctly environment the test (same used in Circle-CI context)
