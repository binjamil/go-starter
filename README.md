# Go Starter

This template can be used as a starting point for developing modern web services in Go.

## Features
- Business logic decoupled via domain in `internal` directory
- Dependency injection employed to make testing easier
- Environment-specific configuration files in `config` directory
- Auto-generated OpenAPI docs by running `make docs`
- [Ent](https://entgo.io/), a rich, type-safe ORM that supports auto-migrations

## Prerequisites

- [Go 1.18](https://golang.org/)
- [Docker 20.10](https://www.docker.com/)

## Running locally

### Set up dependencies

```sh
make compose-up
```

### Run API server

```sh
make run
```

Server will be live now on http://localhost:8080  
Swagger API documentation will be live on http://localhost:8080/swagger/index.html

### Tear down dependencies

```sh
make compose-down
```

## Generate ent code

```sh
make ent
```

Click [here](https://entgo.io/docs/getting-started/) to find more about ent

## Run tests

```sh
make test
```
