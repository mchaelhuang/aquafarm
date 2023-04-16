# Aquafarm

A prototype app that served CRUD to manage farms & ponds. Implementing Clean Architecture to achieve maintainable code.
Ready to run in docker.

## Prerequisite

- go (v1.20)
- Docker
- wire - [github.com/google/wire](https://github.com/google/wire)
- goose - [github.com/pressly/goose](https://github.com/pressly/goose)

## How To Run (in Docker)

1. Run docker compose
   ```shell
   docker-compose up -d
   ```
2. Server will up and listening on `localhost:8080`

## How To Run (in Local)

When runs in local, by default, we still need to use docker to host postgres & redis

1. Copy configuration file in `config/` directory.
   ```shell
   cp config/app.sample.js config/app.development.js
   ```
2. Run docker container
   ```shell
   docker-compose up -d
   ```
3. Build & run binary
   ```shell
   make run
   ```
4. App will be listening on `localhost:8081`

## Development

### Rebuild app container

To reflect changes to docker app container, you can run this following command

```shell
make docker-rebuild-app 
```

### More commands

for more commands you can check available command by run `make help`