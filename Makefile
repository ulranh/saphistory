PACKAGE_NAME := github.com/ulranh/saphistory
GOLANG_CROSS_VERSION ?= v1.16.1

build:
	go build
.PHONY: build

client:
	npm --prefix ./client install
	npm --prefix ./client run build
.PHONY: client

generate-build:
	go generate cmd/assets.go
	go build
.PHONY: generate-build

docker-build:
	docker build --build-arg PORT=${PORT} --build-arg TLS_PATH=${TLS_PATH} -t saphistory:$(VERSION) .
.PHONY: docker-build

# docker build --no-cache -t $(APP_NAME):$(VERSION) .
docker-build-nc:
	docker build --no-cache --build-arg PORT=${PORT} --build-arg TLS_PATH=${TLS_PATH} -t saphistory:$(VERSION) .
.PHONY: docker-build-nc

docker-run:
	docker run -d --name=saphistory -p $(PORT):${PORT} -v $(DBSTORE):/app/badger saphistory:$(VERSION)
.PHONY: docker-run

docker-tag-version:
	docker tag saphistory:$(VERSION) $(DOCKER_REPO)/saphistory:$(VERSION)
.PHONY: docker-tag-version

docker-push-version:
	docker push  $(DOCKER_REPO)/saphistory:$(VERSION)
.PHONY: docker-push-version
