postgresinit:
	docker run --name go-practice -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it go-practice psql

createdb:
	docker exec -it go-practice createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it go-practice dropdb go-chat

.PHONY: postgresinit

