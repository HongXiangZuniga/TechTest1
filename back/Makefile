include .env
export 
run:
	go mod download
	go run cmd/server/api.go

install:
	go mod download

build:
	go build cmd/server/api.go

unit-test:
	go test ./test

docker-build:
	docker build \
	-f build/docker/Dockerfile \
	-t  test-yapo:local .

docker-run:
	docker run --rm -it -p $(PORT):$(PORT) \
	--env-file ./.env \
	test-yapo:local

