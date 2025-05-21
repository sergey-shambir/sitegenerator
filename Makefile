.PHONY: all clean build test image

RELEASE?=0.1.0
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d %H:%M:%S')

all: tidy test build test-js image

clean:
	rm -f bin/sitegenerator

build:
	cd src/sitegenerator && \
	go build -o ../../bin/sitegenerator \
		-ldflags '-s -w -X main.AppVersion=${RELEASE} -X main.AppCommit=${COMMIT} -X "main.AppBuildTime=${BUILD_TIME}"' && \
	cd ../..

test:
	cd src/sitegenerator && \
	go test ./... && \
	cd ../..

test-js:
	cd src/nodeconverter && \
	npm run test && \
	cd ../..

tidy:
	cd src/sitegenerator && \
	go mod tidy && \
	go mod vendor && \
	cd ../..

image: tidy test build
	docker build . -t docker.local/sitegenerator/sitegenerator:latest
