# ${{values.component_id}}

[![CircleCI |<<https://circleci.com/gh/dafiti-group/>${{values.component_id}}/tree/main.svg?style=svg&circle-token=95ef6b5263b43165270e49a6a8a20267dec56af0) |<https://circleci.com/gh/dafiti-group/${{values.component_id}}/tree/main> |

## ${{values.description}}

## How to execute template local

To run the project locally, replace the variables with the tag `${{values.component_id}}` with the name of the desired Application.

## Links devops Tools

| Tool         | URL                                                                                                                                     |
| :----------- | :-------------------------------------------------------------------------------------------------------------------------------------- |
| CicleCI      | <https://app.circleci.com/pipelines/github/dafiti-group/${{values.component_id}>                                                        |
| SonarCube    | <http://sonarqube.live.br.dafiti.io/dashboard?id=${{values.component_id}}>                                                              |
| Argo QA      | <https://argocd.dafiti.la/applications/${{values.component_id}}-br-qa>                                                                  |
| Argo Live    | <https://argocd.dafiti.la/applications/${{values.component_id}}-br-live>                                                                |
| GrayLog QA   | <https://graylog.qa.dafiti.local/search?q=kubernetes_namespace_name%3A%22${{values.component_id}}%22&rangetype=relative&relative=300>   |
| GrayLog Live | <https://graylog.live.dafiti.local/search?q=kubernetes_namespace_name%3A%22${{values.component_id}}%22&rangetype=relative&relative=300> |
| Instana      | <https://apm-dafiti.instana.io/#/services>                                                                                              |

## Hosts

| Env        | Path             | URL                                                                             |
| :--------- | :--------------- | :------------------------------------------------------------------------------ |
| Local      | root             | <http://localhost:8080>                                                         |
| Local      | Liveness Health  | <http://localhost:8080/health-check/liveness>                                   |
| Local      | Readiness Health | <http://localhost:8080/health-check/readiness>                                  |
|            |                  |                                                                                 |
| QA         | root             | <https://${{values.component_id}}.eks.qa.dafiti.local>                          |
| QA         | Liveness Health  | <https://${{values.component_id}}.eks.qa.dafiti.local/health-check/liveness>    |
| QA         | Readiness Health | <https://${{values.component_id}}.eks.qa.dafiti.local/health-check/readiness>   |
|            |                  |                                                                                 |
| Production | root             | <https://${{values.component_id}}.eks.live.dafiti.local>                        |
| Production | Liveness Health  | <https://${{values.component_id}}.eks.live.dafiti.local/health-check/liveness>  |
| Production | Readiness Health | <https://${{values.component_id}}.eks.live.dafiti.local/health-check/readiness> |

## Below is the config you need to do before execute Project

- Be sure about your permissions in [THIS CONFLUENSE PAGE |<https://dafiti.jira.com/wiki/spaces/DFTEC/pages/3247013947/Desenvolvimento+local+DOCKER-DAFITI> |

## Development (outside dev-container VScode)

The requirement to use this repo is `docker` and `docker-compose >=v1.27`, if you need
to install this plugin, follow these steps:

- install [Docker |<https://docs.docker.com/engine/install/ubuntu/> |
- install [Docker compose |<https://docs.docker.com/compose/install/> |

and all you need to run is...

```sh
docker-compose up app
```

If your application is not up and running, try to double check the envs in `docker-compose.yaml` file

## Testing

To test your application you need to run `docker-compose run --rm test`
and this command will use the correctly environment the test (same used in Circle-CI context)
