include .env

POSTGRESQL_URL='postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' 

createdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USERNAME) -c "CREATE DATABASE $(DB_NAME);"

dropdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USERNAME) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

docker_createdb:
	docker exec -it postgres createdb --username=$(DB_USERNAME) --owner=$(DB_USERNAME) $(DB_NAME)

docker_dropdb:
	docker exec -it postgres dropdb $(DB_NAME)

migrateup:
	migrate -path db/migration -database $(POSTGRESQL_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(POSTGRESQL_URL) -verbose down

sqlc:
	sqlc generate

start-crawl:
	cd services/ || go run main.go processDeck.go processTournament.go card_helper.go

getdata:
	make createdb || make migrateup || make start-crawl

.PHONY: postgres createdb dropdb docker_createdb docker_dropdb migrateup migratedown sqlc getdata start-crawl