# ${{values.component_id}}

[![CircleCI](https://circleci.com/gh/dafiti-group/${{values.component_id}}/tree/main.svg?style=svg&circle-token=95ef6b5263b43165270e49a6a8a20267dec56af0)](https://circleci.com/gh/dafiti-group/${{values.component_id}}/tree/main)

## ${{values.description}}

## How to execute template local

Realizar replace da variavel `${{values.component_id}}` pelo nome da Aplicação.

## Hosts

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

---

## Testing

To test your application you need to run `docker-compose run --rm test`
and this command will use the correctly environment the test (same used in Circle-CI context)
