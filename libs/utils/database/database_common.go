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

func NewDatabaseManager(dbName string) (*DatabaseManager, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		dbName)
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

func (dm *DatabaseManager) AddCardToDeck(card int, deck int) (bool, error) {
	ctx := context.Background()
	// TODO: check for the type addcardtodeck return
	fmt.Println("hello")
	_, err := dm.querier.AddCardToDeck(ctx, db.AddCardToDeckParams{
		CardID: sql.NullInt64{Int64: int64(card), Valid: true},
		DeckID: sql.NullInt64{Int64: int64(deck), Valid: true},
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
