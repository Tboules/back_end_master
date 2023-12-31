postgres:
	docker run --name pg_dock -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it pg_dock createdb --username=root --owner=root simple_bank

deletedb:
	docker exec -it pg_dock dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package=mockdb -destination=./db/mock/store.go github.com/Tboules/back_end_master/db/sqlc Store

.PHONY: postgres createdb deletedb migrateup migratedown sqlc test server mock
