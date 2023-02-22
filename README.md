# Golang REST API application for managing images data.

## Note

PostgresSQL should be pre-installed.

Used docker version 20.10.21 and docker-compose 2.12.2

.env file is required with the file name in format 'config.(mode).env' see config.dev.env for example.

## How to run
for development:
```
make run-dev
```
or with docker:
```
make run-dev-docker
```

### For API Specification
run this application and access http://localhost:3000/swagger/

## How to Stop
with docker:
```
make stop-dev-docker
```

## Build this application
```
make build
```
or build for docker image
```
make build-docker
```