BUILD_PARAMS="-buildvcs=false"

.PHONY: all build clean

all: build

build:
	go build "${BUILD_PARAMS}" -v ./cmd/fdb-exporter

clean:
	rm -f ./fdb-exporter
