GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/render

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -rf render render.exe
