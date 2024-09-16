postgres:
    docker run --name postgres -p 15232:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -d postgres-oasis

createdb:
	docker-exec -it postgres createdb --username=username --owner=password oasis

dropdb:
	docker-exec -it postgres dropdb --username=username --owner=password oasis

migrateup:
	migrate -path db/migration -database "postgresql://username:password@localhost:15232/oasis?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://username:password@localhost:15232/oasis?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: postgres createdb dropdb migrateup migratedown