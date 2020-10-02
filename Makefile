VERSION = $(shell git describe --tags)
BUILD_LDFLAGS = "-s -w -X main.Version=${GIT_VER}"

export GO111MODULE=on

run:
	go run main.go

build:
	go build -ldflags=$(BUILD_LDFLAGS)

test:
	go test

packages:
	goxz -pv=$(VERSION) -static -build-ldflags=$(BUILD_LDFLAGS) -d=./pkg/
