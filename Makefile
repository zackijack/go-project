SHELL         = /bin/bash

APP_NAME      = go-project
VERSION      := $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_COMMIT    = $(shell git rev-parse HEAD)
BUILD_DATE    = $(shell date '+%Y-%m-%d-%H:%M:%S')

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Make commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build            Compile the project.'
	@echo '    make run ARGS=        Run with supplied arguments.'
	@echo '    make test             Run tests on a compiled project.'
	@echo '    make clean            Clean the directory tree.'

	@echo

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -X github.com/zackijack/${APP_NAME}/internal/version.Version=${VERSION} -X github.com/zackijack/${APP_NAME}/internal/version.GitCommit=${GIT_COMMIT} -X github.com/zackijack/${APP_NAME}/internal/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

.PHONY: run
run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: test
test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go mod tidy
	go test ./...

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}
