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
	conn    *sql.DB
}

func (dm *DatabaseManager) Close() error {
	return dm.conn.Close()
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
		conn:    conn,
	}, nil
}

func (dm *DatabaseManager) AddCardToDeck(card int, deck int, count int) (bool, error) {
	ctx := context.Background()
	_, err := dm.querier.AddCardToDeck(ctx, db.AddCardToDeckParams{
		CardID:    int64(card),
		DeckID:    int64(deck),
		CardCount: int32(count),
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dm *DatabaseManager) AddCard(data db.CreateCardParams) (bool, error) {
	ctx := context.Background()
	_, err := dm.querier.CreateCard(ctx, data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dm *DatabaseManager) AddDeck(data db.CreateDeckParams) (bool, error) {
	ctx := context.Background()
	_, err := dm.querier.CreateDeck(ctx, data)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (dm *DatabaseManager) AddTournament(data db.CreateTournamentParams) (bool, error) {
	ctx := context.Background()
	_, err := dm.querier.CreateTournament(ctx, data)
	if err != nil {
		return false, err
	}
	return true, nil
}
