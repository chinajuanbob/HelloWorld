GOOS?=linux
GOARCH?=amd64
REPO_LOCAL?=chinajuanbob.com
ROOT_PKG?=github.com/chinajuanbob/HelloWorld
TAG?=v1

.PHONY: all
all: build image run

.PHONY: setup
setup:
	go get github.com/micro/protoc-gen-micro
	go get github.com/go-swagger/go-swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: clientgen
clientgen:
	rm -rf pkg/gen/*
	curl http://localhost:9999/api/v1/swagger.json | jq "." >swagger.json
	swagger generate client -f swagger.json -A todoClient -c pkg/gen/client -m pkg/gen/modules --default-scheme=http --skip-validation
	rm -f swagger.json

.PHONY: codegen
codegen:
	protoc --micro_out=. --go_out=. ./pb/todo.proto

.PHONY: build
build: codegen
	rm -f build/*	
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
	  -installsuffix cgo \
	  -ldflags "-s -w" \
	  -o build/hw \
	  $(ROOT_PKG)/cmd/hw
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/hwcli $(ROOT_PKG)/cmd/hwcli

.PHONY: image
image:
	docker build -t $(REPO_LOCAL)/helloworld-$(GOARCH):$(TAG) .
	# docker push $(REPO_LOCAL)/master-$(GOARCH):$(TAG)

.PHONY: run
run:
	docker run -p6666:6666 $(REPO_LOCAL)/helloworld-$(GOARCH):$(TAG) serve grpc