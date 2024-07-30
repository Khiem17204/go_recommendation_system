package main

import (
	"encoding/json"
	"fmt"
	db "go-rec-sys/db/sqlc"
	utils "go-rec-sys/libs/utils/class"
	database "go-rec-sys/libs/utils/database"
	"io"
	"net/http"
	"strconv"
)

// https://ygoprodeck.com/api/decks/getDeckInfo.php?deckId=506208
type processDeck struct {
	tournament_id string
	deck_id       string
	databaseConn  *database.DatabaseManager
}

func NewProcessDeck(deck_id string, tournament_id string, databaseConn *database.DatabaseManager) *processDeck {
	return &processDeck{
		deck_id:       deck_id,
		tournament_id: tournament_id,
		databaseConn:  databaseConn,
	}
}

// return list of ID of all the cards if success, return nil if failed
func (p *processDeck) getDeck() (*utils.Deck, error) {
	res, err := http.Get("https://ygoprodeck.com/api/decks/getDeckInfo.php?deckId=" + p.deck_id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error:", readErr)
		return nil, readErr
	}
	fmt.Println("Deck ID:", p.deck_id, "successfully fetched")
	var deck utils.Deck
	err = json.Unmarshal(body, &deck)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return &deck, nil
}

// upsert the deck data into the database
func (p *processDeck) upsertCard(deck utils.Deck) (bool, error) {
	deck_id, err := strconv.Atoi(p.deck_id)
	if err != nil {
		fmt.Println("Error: while converting deck ID in mainDeck to int", err)
		return false, err
	}
	frequencyCard := make(map[int]int)
	for _, card := range deck.Maindeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		frequencyCard[card_id]++

	}

	for _, card := range deck.Extradeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		frequencyCard[card_id]++
	}

	for _, card := range deck.Sidedeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		frequencyCard[card_id]++
	}
	// process all cards in the deck
	for card_id, card_count := range frequencyCard {
		_, err := p.databaseConn.AddCardToDeck(card_id, deck_id, card_count)
		if err != nil {
			fmt.Println("Error:", err)
			return false, err
		}
	}
	// return true if successfully add all card to the deck
	return true, nil
}

func (p *processDeck) upsertDeck(deck utils.Deck) (bool, error) {
	rawDeckInfo := [][]string{deck.Maindeck, deck.Extradeck, deck.Sidedeck}
	jsonData, err := json.Marshal(rawDeckInfo)
	if err != nil {
		fmt.Println("Error: while marshalling deck info", err)
		return false, err
	}
	int_tournament_id, err := strconv.Atoi(p.tournament_id)
	if err != nil {
		fmt.Println("Error: while converting tournament ID to int", err)
		return false, err
	}
	data := db.CreateDeckParams{
		ID:           int64(deck.ID),
		DeckName:     deck.DeckName,
		Rank:         deck.Rank,
		TournamentID: int64(int_tournament_id),
		RawDeckInfo:  string(jsonData),
	}
	res, err := p.databaseConn.AddDeck(data)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return res, err
}
