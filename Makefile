DOCKER_REPO ?= docker-lab.repo.theplatform.com

IMAGE_NAME := entarchdemo
IMAGE_VERSION := 0.0.1
IMAGE_NAME_VERSION := $(IMAGE_NAME):$(IMAGE_VERSION)

GITREV = $(shell git rev-parse --short HEAD)
BUILDDATE = $(shell date -u +%Y-%m-%dT%H:%M:%SZ00:00)
LDFLAGS = -X main.VERSION=$(IMAGE_VERSION) -X main.GITREV=$(GITREV) -X main.BUILDDATE=$(BUILDDATE)

default: build

init: clean
	@echo "installing dependencies"
	go get ./...

clean:
	@echo "cleaning"
	go clean -x
	rm -f ./.out/*
#	docker rmi -f ${IMAGE_NAME_VERSION}

build:
	@echo "building"
	go build ./...
	go build -ldflags "$(LDFLAGS)" -o ./.out/${IMAGE_NAME} .

test:
	@echo "running tests"
	go get -t ./...
	go test -v -cover -short ./...

tag:
	docker tag $(IMAGE_NAME) $(DOCKER_REPO)/$(IMAGE_NAME_VERSION)

buildDocker:
	@echo "building docker image"
	docker build --rm --label URCS_SERVICE_VERSION="${IMAGE_VERSION}" -t ${IMAGE_NAME} .

push: buildDocker tag
	@echo "pushing docker image"
	docker push $(DOCKER_REPO)/$(IMAGE_NAME_VERSION)

itest:
	@echo "running integration tests"
	go get -t ./...
	go test -v -tags=integration

# Local build and run in docker
buildLocalForLinux:
	@echo "building for Linux"
	go build ./...
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ./.out/${IMAGE_NAME} .

buildLocalDockerImage:
	@echo "building docker image"
	docker rmi -f ${IMAGE_NAME_VERSION}
	docker build --rm --label URCS_SERVICE_VERSION="${IMAGE_VERSION}" -t ${IMAGE_NAME_VERSION} .

start-docker: init buildLocalForLinux test buildLocalDockerImage
	@echo "running docker image"
#	docker run ${IMAGE_NAME_VERSION}
#   Make sure to have the profile which you are using below is set with correct secret and access keys in your local aws creds file.
	docker run \
		-e AWS_PROFILE=test-account \
		-v ${HOME}/.aws/credentials:/root/.aws/credentials:ro \
		-p 10532:10532 \
		 ${IMAGE_NAME_VERSION}
