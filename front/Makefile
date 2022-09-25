include .env
export 

install:
	npm install

run:
	npm run serve -- --port $(PORT)

build:
	npm run build

lint:
	npm run lint

unit-test:
	npm run test:unit


docker-build-m1:
	docker build \
	-f build/docker/m1/Dockerfile \
	-t test-yapo/front:local .


docker-run:
	docker run --rm -it \
	--env-file ./.env \
	-p $(PORT):80 \
	test-yapo/front:local