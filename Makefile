include .env

postgres:
	docker run --name $(DB_HOST) -p $(DOCKER_DB_PORT):$(DB_PORT) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:16.3-alpine3.20

createdb:
	docker exec -it postgres createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) go-sys-rec

dropdb:
	docker exec -it postgres dropdb go-sys-rec

migrateup:
	migrate -path db/migration -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:$(DB_PORT)/go-sys-rec?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:$(DB_PORT)/go-sys-rec?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc