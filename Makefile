.PHONY: build test gen-spec

PACKAGE_LIST=$(shell go list ./... | grep -v lolapi)

test:
	go test ${PACKAGE_LIST}
lint:
	golangci-lint run -v
clean:
	rm -rf lolapi/*
gen-spec: clean
	swagger generate client -f specs/api.yaml -t lolapi --with-flatten=remove-unused -A LeagueOfLegends
build:
	go get -t -v ./...
	go build .
test-with-coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic ${PACKAGE_LIST}
