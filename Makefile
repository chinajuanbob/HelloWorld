GOOS?=linux
GOARCH?=amd64
REPO_LOCAL?=chinajuanbob.com
ROOT_PKG?=github.com/chinajuanbob/HelloWorld
TAG?=v1

.PHONY: all
all: build image

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