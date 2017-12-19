BINARY = lsgbc
GOARCH = amd64
LDFLAGS = -ldflags "-s -w"

# Installs goalng `dep`, than project dependencies and builds binaries.
all: clean dep ensure linux darwin windows
# Just builds binaries (expect that dependencies are satisfied).
bin: linux darwin windows

dep:
	go get -u github.com/golang/dep/cmd/dep

ensure:
	dep ensure

linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH}

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe

clean:
	-rm -f ${BINARY}-*
	-rm -rf vendor

.PHONY: linux darwin windows clean dep ensure
