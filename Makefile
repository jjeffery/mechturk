.PHONY: \
	build \
	test \
	generate \
	fmt

build: fmt
	go install  ./cmd/...

test:
	go test ./...

generate:
	go generate ./...

fmt:
	go fmt ./...
