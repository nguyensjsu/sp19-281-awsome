
all: clean

clean: 
	find . -name 'search' -type f -exec rm -f {} \;
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt search

install:
	go install search

build:
	go build search

start:
	./search

test-ping:
	curl localhost:3000/ping

test-gumball:
	curl localhost:3000/gumball

docker-build: 
	docker build -t search .
	docker images

docker-run:
	docker run --name search -td -p 3000:3000 search
	docker ps

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it search bash

docker-clean:
	docker stop search
	docker rm search
	docker rmi search

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"

