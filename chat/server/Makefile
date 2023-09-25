postgresinit:
	sudo docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	sudo docker exec -it postgres15 psql

createdb:
	sudo docker exec -it postgres15 createdb --username=root --owner=root go-chat

dropdb:
	sudo docker exec -it postgres15 dropdb go-chat

.PHONY: postgresinit

