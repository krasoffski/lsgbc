
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

tar.gz: clean/dist/tar.gz dist
	tar -zcvf dist/${BINARY}-${VERSION}-linux-${GOARCH}.tar.gz -C dist ${BINARY}-linux-${GOARCH}
	tar -zcvf dist/${BINARY}-${VERSION}-darvin-${GOARCH}.tar.gz -C dist ${BINARY}-darwin-${GOARCH}
	tar -zcvf dist/${BINARY}-${VERSION}-windows-${GOARCH}.tar.gz -C dist ${BINARY}-windows-${GOARCH}.exe
	cd dist && md5sum *.tar.gz > md5sum.txt

zip: clean/dist/zip dist
	cd dist && zip ${BINARY}-${VERSION}-linux-${GOARCH}.zip ${BINARY}-linux-${GOARCH}
	cd dist && zip ${BINARY}-${VERSION}-darvin-${GOARCH}.zip ${BINARY}-darwin-${GOARCH}
	cd dist && zip ${BINARY}-${VERSION}-windows-${GOARCH}.zip ${BINARY}-windows-${GOARCH}.exe
	cd dist && md5sum *.zip > md5sum.txt

clean/dist:
	-rm -rf dist

clean/dist/tar.gz:
	-rm -rf dist/*.tar.gz

clean/dist/zip:
	-rm -rf dist/*.zip

clean/vendor:
	-rm -rf vendor

.PHONY: dep dist clean
