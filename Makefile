all: build

ENVVAR = GOOS=linux CGO_ENABLED=0

build: clean
	$(ENVVAR) go build -o kube-applier

container:
	docker build -t ghcr.io/jrcichra/kube-applier .

clean:
	rm -f kube-applier

test-unit: clean build
	$(GODEP_BIN) go test -v --race ./...

.PHONY: all build container clean test-unit
