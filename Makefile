GOOS?=linux
GOARCH?=amd64
REPO_LOCAL?=chinajuanbob.com
ROOT_PKG?=github.com/chinajuanbob/HelloWorld
TAG?=v1

.PHONY: all
all: build image

.PHONY: setup
setup:
	go get github.com/micro/protoc-gen-micro

.PHONY: codegen
codegen:
	protoc --micro_out=. --go_out=. ./pb/todo.proto

.PHONY: build
build: 
	rm -f build/server	
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
	  -installsuffix cgo \
	  -ldflags "-s -w" \
	  -o build/server \
	  $(ROOT_PKG)/cmd/server

.PHONY: image
image:
	docker build -t $(REPO_LOCAL)/helloworld-$(GOARCH):$(TAG) .
	# docker push $(REPO_LOCAL)/master-$(GOARCH):$(TAG)