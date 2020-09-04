build:
	go build

test: build
	go test -v -count=1 ./testing

PHONY: build test
