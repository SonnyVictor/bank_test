postgres:
	docker run --name simple_bank -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it simple_bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simple_bank dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5000/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5000/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
