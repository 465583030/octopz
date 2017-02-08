deps:
	go get gopkg.in/redis.v5
	go get github.com/kelseyhightower/envconfig
	go get github.com/gorilla/mux

test:
	go test ./...

startredis:
	docker rm -f octopz-redis
	docker run --publish 6379:6379 --name octopz-redis -d redis

run:
	go build ./ && ./octopz
