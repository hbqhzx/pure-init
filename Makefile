.PHONY: all build debug run

BUILD_IMAGE ?= hub.ark.jd.com/chenyao6/golang:1.9_jenkins_root
PRODUCT_IMAGE ?= 192.168.104.19:5000/sky/deploy:v0.1
VET_IMAGE ?= 192.168.104.19:5000/hawkeye/centos6-golang:go1.9

CWD := $(shell pwd)
CONTAINER_NAME := deploy-compile

JSS_PROFILE ?= ark
BUCKET ?= skywing-deploy-build
LATEST_PKG ?="skywing-deploy-build_latest.tar.gz"
INSTALLER := output/target/sky-deploy.tar.gz
export GOPATH=$(shell pwd)

DOCKER_RUN_OPTS := \
	--name $(CONTAINER_NAME) \
    --rm \
	--hostname $(CONTAINER_NAME) \
	--volume=/etc/localtime:/etc/localtime:ro \
	--volume=$(CWD):$(CWD) \
	--workdir=$(CWD)

DOCKER_VET_RUN_OPTS := \
        -t \
        --rm=true \
        --publish-all=true \
        --privileged=true \
        --userns=host   \
        --name $(CONTAINER_NAME) \
        --hostname $(CONTAINER_NAME) \
        --volume=/etc/localtime:/etc/localtime:ro \
        --volume=$(CWD):/ws \
        --workdir=/ws  \


all: build

vet:
	docker run $(DOCKER_VET_RUN_OPTS) $(VET_IMAGE) ./go_vet.sh

compile:
	if docker ps -a | grep -q $(CONTAINER_NAME); then \
            docker start $(CONTAINER_NAME); \
            docker exec $(CONTAINER_NAME) ./build.sh; \
        else \
            docker run $(DOCKER_RUN_OPTS) $(BUILD_IMAGE) ./build.sh; \
        fi
build:
	./build.sh -ci

debug:
	if docker ps -a | grep -q $(CONTAINER_NAME); then \
            docker start $(CONTAINER_NAME); \
            docker exec -it $(CONTAINER_NAME) /bin/bash; \
        else \
            docker run -it $(DOCKER_RUN_OPTS) $(BUILD_IMAGE) /bin/bash; \
        fi

image:
	docker build -t $(PRODUCT_IMAGE) $(CWD)
	# docker push $(PRODUCT_IMAGE)

upload:
	jsutil jss cp $(INSTALLER) jss://$(BUCKET)/$(LATEST_PKG) --profile $(JSS_PROFILE)

run:
	go install api && ./bin/api
