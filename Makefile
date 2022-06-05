createdb:
	docker exec -it blog-backend createdb --username=root blog

dropdb:
	docker exec -it blog-backend dropdb --username=root blog

createschema:
	migrate create -ext sql -dir migration -seq init_schema

migrateup:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose down

migrateup1:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path migration -database "postgresql://root:secret@localhost:5432/blog?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go1.18 run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/teguhatma/simple-bank/db/sqlc Store

.PHONY: createdb dropdb createschema migratedown migrateup sqlc test server mock migratedown1 migrateup1