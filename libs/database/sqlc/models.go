// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Card struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	FrameType   string          `json:"frame_type"`
	Archetype   sql.NullString  `json:"archetype"`
	Attribute   sql.NullString  `json:"attribute"`
	Race        sql.NullString  `json:"race"`
	Level       sql.NullInt32   `json:"level"`
	Attack      sql.NullInt32   `json:"attack"`
	Defense     sql.NullInt32   `json:"defense"`
	Description string          `json:"description"`
	RawCardInfo json.RawMessage `json:"raw_card_info"`
}

type CardsInDeck struct {
	CardID    int64 `json:"card_id"`
	DeckID    int64 `json:"deck_id"`
	CardCount int32 `json:"card_count"`
}

type Deck struct {
	ID           int64  `json:"id"`
	DeckName     string `json:"deck_name"`
	Rank         string `json:"rank"`
	TournamentID int64  `json:"tournament_id"`
	RawDeckInfo  string `json:"raw_deck_info"`
}

type Tournament struct {
	ID                int64         `json:"id"`
	TournamentName    string        `json:"tournament_name"`
	Tier              int32         `json:"tier"`
	PlayerCount       sql.NullInt32 `json:"player_count"`
	EventDate         time.Time     `json:"event_date"`
	Format            string        `json:"format"`
	RawTournamentInfo string        `json:"raw_tournament_info"`
}
