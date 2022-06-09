# Golang Template Project

___

*Issue to build local image to develop because docker `dafiti-certificates` image?*

*Use this [Confluense page](https://dafiti.jira.com/l/c/C6Y6BHjW) to guide you to fix it*

___

[comment]: <!-- (You need to change all `golang-template-project`to your repo name) -->

> This microservice is part of Dafiti-Group project template `golang-template-project`
and has the objective to deliver **HERE THE PURPOSE OF THIS MICROSERVICE**

You are brand new here? so read these files:

1. How do I use this repository as base project? Follow the steps [HERE](./docs/USING-THIS-PROJECT-AS-BASE.md)
1. How to configure your new microservice in `argo and charts` [HERE](./docs/CONFIG-CHART-AND-ARGO.md)
1. How to update `terraform` templates in `OPS` [HERE](./docs/CONFIG-OPS.md)
1. How to validate `circle ci` steps [HERE](./docs/CONFIG-CICD.md)
1. How to adjust my application configurations [HERE](./docs/CONFIG-MS.md)
1. Using VScode devcontainer to development [HERE](./docs/VSCODE.md)

___

## Microservice chart mapping

This template project does not have an chart deployed in
[Dafiti Group template repo](https://github.com/dafiti-group/diagrams)
as an exemple, but you must create a draw with the integrations of
your new ms to help team workers to understand the meaning of this
application.

Plese, take a time and read this [DOCUMENTATION](./docs/CONFIG-DIAGRAMS.md)

___

## Argo Deployer: status by the environment

> The badges information to update the links below are placed in Argo application information

- Dashboard: [https://argocd.dafiti.la](https://argocd.dafiti.la) and login with google SSO
- live:
  - [![LIVE Build Status](https://argocd.dafiti.la/api/badge?name=golang-template-project-br-live&revision=true)](https://argocd.dafiti.la/api/badge?name=golang-template-project-br-live&revision=true)
- QA:
  - [![QA Build Status](https://argocd.dafiti.la/api/badge?name=golang-template-project-br-qa&revision=true)](https://argocd.dafiti.la/api/badge?name=golang-template-project-br-qa&revision=true)

___

## Configurations

The configuration of this project is not present in this repository.
This information is split into two different places as below:
[Charts](https://github.com/dafiti-group/charts)
and
[Argo](https://github.com/dafshowediti-group/argo) repositories.

*__Cluster resources and limits__* are placed in
[Charts repo in values.yaml file](https://github.com/dafiti-group/charts/blob/master/charts/golang-template-project/values.yaml)

*__Live environmnet__* are placed in
[argo golang-template-project.yaml](https://github.com/dafiti-group/argo/blob/master/clusters/eks-live-dafiti-latam/apps/golang-template-project.yaml)

*__QA environmnet__* are placed in
[argo golang-template-project.yaml](https://github.com/dafiti-group/argo/blob/master/clusters/eks-qa-dafiti-latam/apps/golang-template-project.yaml)

___

## Deploy

The deployment is automatic by `Circle-CI` so, the flow of this project is:

- all feature branches will be deployed to QA
- Only branch master will deploy to QA (staging) and, with approval, deploy to live

### Available hosts

- live:
  - <https://golang-template-project.eks.pub.live.dafiti.io>
  - <https://golang-template-project.eks.priv.live.dafiti.local>
- qa:
  - <https://golang-template-project.eks.pub.qa.dafiti.io>
  - <https://golang-template-project.eks.priv.qa.dafiti.local>

### Helper URLs

- Newrelic (TODO)
- kibana (TODO)
- graylog (TODO)
- grafana (TODO)
- instana (TODO)

___

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

___

## Testing

To test your application you need to run `docker-compose run --rm test`
and this command will use the correctly environment the test (same used in Circle-CI context)
