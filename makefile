# Go parameters
GOCMD=go
GOGENERATE=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=coda-world
BINARY_UNIX=$(BINARY_NAME)-linux-amd64

build:
	$(GOGENERATE) ./...
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
test:
	$(GOTEST) ./...
clean:
	$(GOCLEAN)
	rm -rdf bin
run:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v
	bin/$(BINARY_NAME)
deps:
	$(GOGET) golang.org/x/tools/cmd/stringer
	$(GOGET) github.com/golang/protobuf/protoc-gen-go

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_UNIX) -v
