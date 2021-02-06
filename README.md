# YOUR APPLICATION NAME

[comment]: <!-- (You need to change all `golang-template-project`to your repo name) -->

> This microservice is part of Dafiti-Group project template `golang-template-project`
and has the objective to deliver *_HERE THE PURPOSE OF THIS MICROSERVICE_*

You are brand new here? so read these files:

1. How to configure your new microservice in `argo and charts` [HERE](./docs/CONFIG-CHART-AND-ARGO.md)
1. How to update `terraform` templates in `OPS` [HERE](./docs/CONFIG-OPS.md)
1. How to validate `circle ci` steps [HERE](./docs/CONFIG-CICD.md)
1. Using vscode devcontainer to development [HERE](./docs/VSCODE.md)

___

## Argo Deployer: status by the environment

> The badges information to update the links below are placed in Argo application information

- Dashboard: [https://argocd.dafiti.la](https://argocd.dafiti.la) and login with google SSO
- live:
  - [![LIVE Build Status]()]()
- QA:
  - [![QA Build Status]()]()

___

## Configurations

The configuration of this project is not present in this repository.
This information is split into two different places as below:
[Charts](https://github.com/dafiti-group/charts)
and
[Argo](https://github.com/dafshowediti-group/argo) repositories.



*__Cluster resources and limits__* are placed in
[Charts repo in values.yaml file](https://github.com/dafiti-group/charts/blob/master/charts/APPLICATION-NAME/values.yaml)

*__Live environmnet__* are placed in
[argo APPLICATION-NAME.yaml](https://github.com/dafiti-group/argo/blob/master/clusters/eks-live-dafiti-latam/apps/APPLICATION-NAME.yaml)

*__QA environmnet__* are placed in
[argo APPLICATION-NAME.yaml](https://github.com/dafiti-group/argo/blob/master/clusters/eks-qa-dafiti-latam/apps/APPLICATION-NAME.yaml)

___

## Deploy

The deployment is automatic by `Circle-CI` so, the flow of this project is:

- all feature branches will be deployed to QA
- Only branch master will deploy to QA (staging) and, with approval, deploy to live

### Available hosts

- live:
  - https://APPLICATION-NAME.eks.pub.live.dafiti.io
  - https://APPLICATION-NAME.eks.priv.live.dafiti.local
- qa:
  - https://APPLICATION-NAME.eks.pub.qa.dafiti.io
  - https://APPLICATION-NAME.eks.priv.qa.dafiti.local

### Helper URLs

- Newrelic (TODO)
- kibana (TODO)
- graylog (TODO)
- grafana (TODO)
- instana (TODO)

___

## Development

The requirement to use this repo is `docker` and `docker-compose >=v1.27`, if you need
to install this plugin, follow this steps:

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
