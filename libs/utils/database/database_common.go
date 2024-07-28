package database

import (
	"context"
	"database/sql"
	"fmt"
	db "go-rec-sys/db/sqlc"
	"os"

	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	querier db.Querier
}

func NewDatabaseManager() (*DatabaseManager, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &DatabaseManager{
		querier: db.New(conn),
	}, nil
}

func (dm *DatabaseManager) AddCardToDeck() (bool, error) {
	ctx := context.Background()
	defer dm.querier.AddCardToDeck(ctx, db.AddCardToDeckParams{})
	return true, nil
}
