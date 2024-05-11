FROM golang:alpine3.19 AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o reservation ./main.go

# final stage
FROM alpine:3.19.1

WORKDIR /app

COPY --from=build-env /build/reservation /app/

ENTRYPOINT ["./reservation"]

CMD [ "-h"]