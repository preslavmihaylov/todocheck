version?=$(shell date -u +'custom_%y-%m-%dT%H:%MZ')
RELEASE_BINARIES=binaries-$(version)

build:
	go build -ldflags "-X main.version=$(version)"

build-test:
	go build -ldflags "-X main.version=$(version) -X main.runningTests=true"

test: build-test
	go test -v -count=1 ./testing

release:
	@echo "Generating binaries for version $(version)..."
	./release.sh $(RELEASE_BINARIES) $(version)
	@echo "Version binaries available in folder $(RELEASE_BINARIES)"

.PHONY: build test release
