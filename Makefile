prepare:
	go mod download

run:
	go build -o bin/main cmd/media/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-delete-media .

dkr:
	docker run -p "7200:7200" -p "8150:8150" superhero-delete-media

launch: dkb dkr

delete-media-log:
	docker logs superhero-delete-media -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

delete-media-ssh:
	docker exec -it superhero-delete-media /bin/bash

PHONY: prepare build dkb dkr launch delete-media-log delete-media-ssh rmc rmi clear