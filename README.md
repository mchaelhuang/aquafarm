# Aquafarm

A prototype app that served CRUD to manage farms & ponds. Implementing Clean Architecture to achieve maintainable code.
Ready to run in docker.

## Prerequisite

- `go v1.20`
- Docker
- `wire` - [github.com/google/wire](https://github.com/google/wire)
- `goose` - [github.com/pressly/goose](https://github.com/pressly/goose)
- `gomock` - [github.com/golang/mock](https://github.com/golang/mock)

## How to Run

To run the app you just need to make sure Docker is installed. The library can be installed later for development.

## Run in Docker

1. Run docker compose

   ```shell
   docker-compose up -d
   ```

2. Server will up and listening on `localhost:8080`

## Run in Local

1. Copy configuration file in `config/` directory.

   ```shell
   cp ./config/app.sample.json ./config/app.development.json
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

Make sure all library listed on prerequisite is installed.
We are using `wire` for dependency injection, `goose` for database migration,
and `gomock` for mocking in unit test.

### Rebuild app container

To reflect changes into docker app container, you can run this following command:

```shell
make docker-rebuild-app 
```

### More commands

You can found out more commands on `make help` command.

## REST Endpoints

For complete documentation you can find it on https://documenter.getpostman.com/view/893849/2s93XyU3JV

```
GET /stats/

GET /v1/farm
GET /v1/farm/:id
POST /v1/farm
PUT /v1/farm/:id
DELETE /v1/farm/:id

GET /v1/pond
GET /v1/pond/:id
POST /v1/pond
PUT /v1/pond/:id
DELETE /v1/pond/:id
```
