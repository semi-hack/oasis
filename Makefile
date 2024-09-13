postgres:
    docker run --name postgres -p 15332:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres-oasis


createdb:
	docker-exec -it postgres createdb --username=root --owner=root oasis

.PHONY: postgres createdb