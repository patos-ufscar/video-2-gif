export GO_VERSION=1.23.3

image=gif-service
version=$(shell git rev-parse --short HEAD)


## Build the service image
.PHONY: image
image:
	docker build . \
		--build-arg GO_VERSION=$(GO_VERSION) \
		-t $(image)

## builds the service
.PHONY: service
service:
	go build -o ./cmd/gif/gif ./cmd/gif

## runs the service locally
.PHONY: run
run: service
	./cmd/gif/gif
