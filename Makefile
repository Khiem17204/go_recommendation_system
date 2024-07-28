include .env

# postgresurl: change url to update database 
POSTGRESQL_URL='postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable' 

createdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USERNAME) -c "CREATE DATABASE $(DB_NAME);"

dropdb:
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USERNAME) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

migrateup:
	migrate -path db/migration -database $(POSTGRESQL_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(POSTGRESQL_URL) -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc