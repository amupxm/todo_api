createdb:
	docker exec -it postgresdb createdb --username=root --owner=root todosdb
dropdb: 
	docker exec -it postgresdb dropdb todosdb
migrate:
	migrate -path db/migrations -database "postgres://root:1234@localhost:5432/todosdb?sslmode=disable" -verbose up