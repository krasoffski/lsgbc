
BINARY    := lsgbc
GOARCH    := amd64
OUTDIR    := bin
WINDOWS   := windows
PLATFORMS := linux darwin windows
ARGUMENTS := --compact --best
COMMIT    ?= $(shell git rev-parse --short HEAD)
TAG       ?= $(shell git describe --tags --dirty 2>/dev/null)
VERSION   := $(if $(TAG),$(TAG),git+$(COMMIT))
LDFLAGS   := -ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY: all
all: clean install

.PHONY: install
install:
	go build ${LDFLAGS} -o ${OUTDIR}/${BINARY}

.PHONY: release
release: clean $(PLATFORMS)

.PHONY: $(PLATFORMS)
windows: SUFFIX=.exe
$(PLATFORMS):
	GOOS=$@ GOARCH=${GOARCH} go build ${LDFLAGS} -o ${OUTDIR}/${BINARY}-$@-${GOARCH}${SUFFIX}
	cd ${OUTDIR} && zip ${BINARY}-${VERSION}-$@-${GOARCH}.zip ${BINARY}-$@-${GOARCH}${SUFFIX}
	cd ${OUTDIR} && md5sum ${BINARY}-${VERSION}-$@-${GOARCH}.zip >> md5sum.txt

.PHONY: run
run:
	@${OUTDIR}/${BINARY} ${ARGUMENTS}

.PHONY: clean
clean:
	-rm -rf ${OUTDIR}
