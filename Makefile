BINARY = lsgbc
GOARCH = amd64
LDFLAGS = -ldflags "-s -w"

# Build the project for platforms
all: clean linux darwin windows

linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH}

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH}

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe

clean:
	-rm -f ${BINARY}-*

.PHONY: linux darwin windows clean
