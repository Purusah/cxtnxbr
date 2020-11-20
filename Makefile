SHELL=/bin/bash

PROJECT=$(shell basename "$(PWD)")
ROOT = $(shell pwd)
API_PORT = 8080
REDIS_PORT=6379
GOLINT_VERSION = v1.31.0

.PHONY: build test web web-build

build:
	docker build --tag $(PROJECT):latest --network "host" --rm --file $(PWD)/build/docker/Dockerfile.dev.app .

lint:
	docker run -t --rm -v $(shell pwd):/$(PROJECT) -w /$(PROJECT) golangci/golangci-lint:$(GOLINT_VERSION) golangci-lint run -v

fmt:
	docker run -it --rm -v $(shell pwd):/$(PROJECT) -w /$(PROJECT) golang:1.15.5-alpine3.12 go fmt -x ./...

run: redis
	docker run -it --rm --name $(PROJECT) \
		--network "host" \
		--env-file scripts/default.env \
		$(PROJECT):latest \
		/app

test: redis
	docker run -it --rm \
		--name $(PROJECT)-test \
		-v $(shell pwd):/$(PROJECT) \
		-w /$(PROJECT) \
		golang:1.15.5-buster \
		go test -v -race -tags=all -coverprofile cp.out ./...
	docker stop $(PROJECT)-redis || true

# UTILS

redis:
	docker stop $(PROJECT)-redis || true
	docker run \
		--rm \
		--name $(PROJECT)-redis \
		--detach \
		--network "host" \
		redis:6
	docker exec -it $(PROJECT)-redis redis-cli set bind 0.0.0.0

redis-cli:
	docker exec -it $(PROJECT)-redis redis-cli
