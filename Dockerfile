FROM golang:alpine3.14 AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o reservation ./main.go

# final stage
FROM alpine:3.13.6

WORKDIR /app

COPY --from=build-env /build/reservation /app/

ENTRYPOINT ./reservation