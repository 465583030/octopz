build:
	docker build -t firemango/octopz .

push:
	docker push firemango/octopz:latest

package:
	GOOS=linux CGO_ENABLED=0 go build -o octopz ../
