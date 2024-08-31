BIN := build/basic-web

ENVS := CGO_ENABLED=0 GOOS=linux
FLAGS := -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo


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

make-profile: 
	curl --output profile.out "localhost:6060/debug/pprof/profile?seconds=30"

make-heap: 
	curl "http://127.0.0.1:55854/debug/pprof/heap?seconds=15" > heap

run-heap:
	go tool pprof -http=:3435 heap

run-profile:
	go tool pprof -http=:3435 pprofile