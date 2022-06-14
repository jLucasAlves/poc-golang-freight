# template-golang

[![CircleCI](https://circleci.com/gh/dafiti-group/template-golang/tree/main.svg?style=svg&circle-token=95ef6b5263b43165270e49a6a8a20267dec56af0)](https://circleci.com/gh/dafiti-group/template-golang/tree/main)

## ${{values.description}}

## How to execute template local

To run the project locally, replace the variables with the tag `template-golang` with the name of the desired Application.

## Hosts

- Local: http://localhost:8080
  - Liveness Health: http://localhost:8080/health-check/liveness
  - Readiness Health: http://localhost:8080/health-check/readiness
- QA: https://template-golang.eks.qa.dafiti.local
  - Liveness Health: https://template-golang.eks.qa.dafiti.local/health-check/liveness
  - Readiness Health: https://template-golang.eks.qa.dafiti.local/health-check/readiness
- Production: https://template-golang.eks.live.dafiti.local
  - Liveness Health: https://template-golang.eks.live.dafiti.local/health-check/liveness
  - Readiness Health: https://template-golang.eks.live.dafiti.local/health-check/readiness

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
