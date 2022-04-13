all: build

ENVVAR = GOOS=linux GOARCH=amd64 CGO_ENABLED=0

build: clean fmt
	$(ENVVAR) go build -o kube-applier

container:
	docker build -t ghcr.io/jrcichra/kube-applier .

clean:
	rm -f kube-applier

fmt:
	find . -path ./vendor -prune -o -name '*.go' -print | xargs -L 1 -I % gofmt -s -w %

test-unit: clean fmt build
	$(GODEP_BIN) go test -v --race ./...

.PHONY: all build container clean fmt test-unit
