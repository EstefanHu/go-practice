build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./...

postgresinit:
	sudo docker run --name go-practice -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	sudo docker exec -it go-practice psql

createdb:
	sudo docker exec -it go-practice createdb --username=root --owner=root go-bank

dropdb:
	sudo docker exec -it go-practice dropdb go-bank

.PHONY: postgresinit

