GOOS?=linux
GOARCH?=amd64
REPO_LOCAL?=chinajuanbob.com
ROOT_PKG?=github.com/chinajuanbob/HelloWorld
TAG?=v1

OPENSHIFT_VERSION=v3.11.0
ISTIO_VERSION=1.1.2
KUBELESS_VERSION=v1.0.3

.PHONY: all
all: build image run

.PHONY: setup
setup:
	go get github.com/micro/protoc-gen-micro
	go get github.com/go-swagger/go-swagger
	go install github.com/go-swagger/go-swagger/cmd/swagger
	brew install openshift-cli
	# brew install source-to-image
	brew install kubernetes-helm
	curl -L https://git.io/getLatestIstio | ISTIO_VERSION=$(ISTIO_VERSION) sh -
	cp istio-$(ISTIO_VERSION)/bin/istioctl /usr/local/bin

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
	cp build/hw build/goapp

.PHONY: image
image:
	docker build -t $(REPO_LOCAL)/helloworld-$(GOARCH):$(TAG) .
	# docker push $(REPO_LOCAL)/master-$(GOARCH):$(TAG)

.PHONY: run
run:
	docker run -p6666:6666 $(REPO_LOCAL)/helloworld-$(GOARCH):$(TAG) serve grpc

.PHONY: start
start:
	minishift profile set servicemesh # TODO: change the name
	minishift config set memory 8GB
	minishift config set cpus 4
	minishift config set image-caching true
	minishift config set openshift-version $(OPENSHIFT_VERSION)
	minishift addon enable admin-user
	minishift addon enable anyuid # TODO: how about normal openshift cluster
	minishift start --vm-driver hyperkit

.PHONY: stop
stop:
	minishift stop

.PHONY: start2
start2:
	oc login https://$$(minishift ip):8443 -u admin -p admin
	# admin
	oc adm policy add-cluster-role-to-group sudoer system:authenticated admin
	# helm
	helm init
	# istio
	kubectl create namespace istio-system
	oc adm policy add-scc-to-user anyuid -z istio-ingress-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z default -n istio-system
	oc adm policy add-scc-to-user anyuid -z prometheus -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-egressgateway-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-citadel-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-ingressgateway-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-cleanup-old-ca-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-mixer-post-install-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-mixer-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-pilot-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-sidecar-injector-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-galley-service-account -n istio-system
	oc adm policy add-scc-to-user anyuid -z istio-security-post-install-account -n istio-system
	# oc adm policy add-scc-to-user anyuid -z kiali-service-account -n istio-system
	helm template istio-$(ISTIO_VERSION)/install/kubernetes/helm/istio-init --name istio-init --namespace istio-system | kubectl apply -f -
	kubectl get crds --namespace=istio-system | grep 'istio.io\|certmanager.k8s.io' | wc -l
	helm template istio-$(ISTIO_VERSION)/install/kubernetes/helm/istio --name istio --namespace istio-system --values istio-$(ISTIO_VERSION)/install/kubernetes/helm/istio/values-istio-demo.yaml | kubectl apply -f -
	kubectl apply -f deploy/clusterrole-kiali.yaml
	# kubeless
	oc create ns kubeless
	oc create -f https://github.com/kubeless/kubeless/releases/download/$(KUBELESS_VERSION)/kubeless-openshift-$(KUBELESS_VERSION).yaml
	# golang s2i
	oc project openshift
	oc import-image chinajuanbob/golang-s2i -n openshift --confirm --insecure
	oc patch is golang-s2i -n openshift -p '{"spec":{"tags":[{"name":"latest","annotations":{"tags":"builder, golang","version":"1.12"}}]}}'

.PHONY: deploy
deploy:
	# create service
	# kubectl apply -f <(istioctl kube-inject -f istio-$(ISTIO_VERSION)/samples/bookinfo/platform/kube/bookinfo.yaml)
	istioctl kube-inject -f istio-$(ISTIO_VERSION)/samples/bookinfo/platform/kube/bookinfo.yaml | oc create -f -
	# manual fix
	# https://github.com/istio/istio/issues/12948

.PHONY: destroy
destroy:
	echo destroy

.PHONY: delete
delete:
	echo delete
