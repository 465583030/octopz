deps:
	go get gopkg.in/redis.v5
	go get github.com/kelseyhightower/envconfig
	go get github.com/gorilla/mux
	
test:

run:
	go build ./ && ./octopz