deps:
	go get gopkg.in/redis.v5
	go get github.com/kelseyhightower/envconfig
	
test:

run:
	go build ./ && ./octopz