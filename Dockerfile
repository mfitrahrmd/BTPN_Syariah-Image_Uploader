FROM golang:alpine AS builder

LABEL maintaner="Muhamad Fitrah Ramadhan <mfitrahrmd>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download
RUN go get -v ./...
RUN go mod tidy

COPY . .

RUN go build -o build/ cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/build/main .
COPY --from=builder /app/*.env .

ENTRYPOINT ["./main"]