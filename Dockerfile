FROM golang:alpine AS builder

LABEL maintaner="Muhamad Fitrah Ramadhan <mfitrahrmd>"

RUN apk update && apk add --no-cache git

WORKDIR /root

COPY . .

RUN go mod download
RUN go get -v ./...
RUN go mod tidy

COPY . .

RUN go build -o build/ cmd/main.go

FROM alpine:latest

WORKDIR /root

COPY --from=builder /root/build/main .
COPY --from=builder /root/*.env .
COPY --from=builder /root/openapi/ ./openapi

EXPOSE 3000

ENTRYPOINT ["./main"]