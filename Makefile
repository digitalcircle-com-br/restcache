GIT_COMMIT := $(shell git rev-list -1 HEAD)
DT := $(shell date +%Y.%m.%d.%H%M%S)
ME := $(shell whoami)
HOST := $(shell hostname)
PRODUCT := restcache


ver:
	echo package main > ver.go
	echo const VER=\"$(DT)\" >> ver.go

run:
	REDIS=redis://localhost:6379 \
	CGO_ENABLED=0 go run -ldflags "-X github.com/digitalcircle-com-br/buildinfo.Ver=$(GIT_COMMIT) -X github.com/digitalcircle-com-br/buildinfo.BuildDate=$(DT) -X github.com/digitalcircle-com-br/buildinfo.BuildUser=$(ME) -X github.com/digitalcircle-com-br/buildinfo.BuildHost=$(HOST) -X github.com/digitalcircle-com-br/buildinfo.Product=$(PRODUCT)" ./main.go

docker_push: docker
	docker tag $(PRODUCT):latest digitalcircle/$(PRODUCT):latest && \
	docker push digitalcircle/$(PRODUCT):latest

docker: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deploy/api -ldflags "-s -w -X github.com/digitalcircle-com-br/buildinfo.Ver=$(GIT_COMMIT) -X github.com/digitalcircle-com-br/buildinfo.BuildDate=$(DT) -X github.com/digitalcircle-com-br/buildinfo.BuildUser=$(ME) -X github.com/digitalcircle-com-br/buildinfo.BuildHost=$(HOST) -X github.com/digitalcircle-com-br/buildinfo.Product=$(PRODUCT)" ./cmd/api/main.go
	cd deploy && \
	docker build -t $(PRODUCT) .

docker_run:
	docker run --rm -it -p 8080:8080 $(AWS_ECR)/$(PRODUCT):latest

