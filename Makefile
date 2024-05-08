NAME?=reservation
ARCH=amd64
BIN = bin/reservation
BIN_LINUX = $(BIN)-linux-$(ARCH)
BIN_DARWIN = $(BIN)-darwin-$(ARCH)
GO111MODULE=on
GIT_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
IMG_NAMESPACE?=docker.io/hriships
IMG_TAG?=v0.0.1
REGISTRY?=$(IMG_NAMESPACE)/$(NAME)

.PHONEY: unit-test all clean

all: $(BIN_LINUX) $(BIN_DARWIN)

$(BIN_DARWIN):
	GOARCH=$(ARCH) GOOS=darwin CGO_ENABLED=0 go build -o $(BIN_DARWIN) main.go

$(BIN_LINUX):
	GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -o $(BIN_LINUX) main.go

unit-test:
	go test -race -count=1 -mod=vendor ./internal/...

integration-test:
	go test -mod=vendor -count=1 ./tests -v

docker: Dockerfile
	docker image build -t "$(REGISTRY):$(IMG_TAG)" .

clean:
	rm -rf bin/