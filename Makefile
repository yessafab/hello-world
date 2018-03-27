run: build start

build:
	go test -v
	docker build . -t hello-world
	@docker rmi -f $$(docker images -f "dangling=true" -q)

start:
	docker run --rm -p 3000:3333 hello-world

curl:
	curl http://
