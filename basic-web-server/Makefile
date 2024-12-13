BIN := build/basic-web

ENVS := CGO_ENABLED=0 GOOS=linux
FLAGS := -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo -gcflags=-m


tidy:
	$(ENVS) go mod tidy -v


.PHONY: build
build: tidy
	mkdir -p build/
	$(ENVS) go build ${FLAGS} -o $(BIN) .

docker-tag:
	docker build -t kadekchresna/basic-web .

docker-push: docker-tag
	docker push kadekchresna/basic-web:latest

docker-run:
	docker run -d --env-file .env kadekchresna/basic-web:latest

profile: 
	curl -o profile.out "http://127.0.0.1:50218/debug/pprof/profile?seconds=30"

heap: 
	curl -o heap.out localhost:8080/debug/pprof/heap?seconds=30 && \
	go tool pprof -dot heap.out > heap.dot && \
	dot -Tpng heap.dot -o heap.png

run-heap:
	go tool pprof -http=:3435 heap

run-profile:
	go tool pprof -http=:3435 pprofile

run-postgres:
	docker run -d \
	--name basic-postgres \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=root \
	-e POSTGRES_DB=postgres \
	-v /Users/kuro/project/golang/go-playground/basic-web-server/deploy/data:/var/lib/postgresql/data \
	-p 5432:5432 \
	postgres:latest

