
TAG?=$(shell git describe --tags --dirty 2>/dev/null)
COMMIT?=$(shell git rev-parse --short HEAD)

VERSION := $(if $(TAG),$(TAG),git+$(COMMIT))

BINARY = lsgbc
GOARCH = amd64
LDFLAGS = -ldflags "-s -w -X main.Version=$(VERSION)"

all: clean dep dist
clean: clean/dist clean/vendor
dep: dep/install dep/ensure
dist: dist/linux dist/darwin dist/windows


dep/install:
	go get -u github.com/golang/dep/cmd/dep

dep/ensure:
	dep ensure

dist/linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-linux-${GOARCH}

dist/linux/short:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}

dist/darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-darwin-${GOARCH}

dist/darwin/short:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}

dist/windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-windows-${GOARCH}.exe

dist/windows/short:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}


clean/dist:
	-rm -rf dist

clean/vendor:
	-rm -rf vendor

.PHONY: dep dist clean
