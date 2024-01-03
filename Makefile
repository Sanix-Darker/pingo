.PHONY: all build clean test deps run docker-build install docker-run compose lint

GO := go
GOFLAGS :=
BINARY_NAME := pingo
DOCKER_IMAGE := pingo:latest

all: clean build run

build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME)

clean:
	$(GO) clean
	rm -rf $(BINARY_NAME)
	docker-compose down
	# docker volume rm pingo_dbdata

install:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

test:
	$(GO) test $(GOFLAGS) ./...

deps:
	$(GO) mod download

run: build
	./$(BINARY_NAME)

lint:
	golangci-lint run

docker-build:
	docker build --rm -t $(DOCKER_IMAGE) -f ./Dockerfile .

docker-run:
	docker run -ti $(DOCKER_IMAGE)

compose: clean
	docker-compose up --build