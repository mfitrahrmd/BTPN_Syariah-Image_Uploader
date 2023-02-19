run-dev:
	APPENV=dev go run cmd/main.go

run-dev-docker:
	docker compose -f dev.docker-compose.yaml up -d

stop-dev-docker:
	docker compose -f dev.docker-compose.yaml down

build:
	go get -v ./... && go mod tidy && go build -o build/ cmd/main.go

build-docker:
	docker build .