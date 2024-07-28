package main

import (
	"encoding/json"
	"fmt"
	utils "go-rec-sys/libs/utils/class"
	database "go-rec-sys/libs/utils/database"
	"io"
	"net/http"
	"strconv"
)

// https://ygoprodeck.com/api/decks/getDeckInfo.php?deckId=506208
type processDeck struct {
	deck_id      string
	databaseConn *database.DatabaseManager
}

func NewProcessDeck(id string) *processDeck {
	databaseConn, err := database.NewDatabaseManager("go_rec_sys")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &processDeck{
		deck_id:      id,
		databaseConn: databaseConn,
	}
}

// return list of ID of all the cards if success, return nil if failed
func (p *processDeck) getCards() (*utils.Deck, error) {
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
func (p *processDeck) upsert() (bool, error) {
	deck, err := p.getCards()
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	// upsert operation, maindeck
	for _, card := range deck.Maindeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		deck_id, err := strconv.Atoi(p.deck_id)
		if err != nil {
			fmt.Println("Error: while converting deck ID in mainDeck to int", err)
			return false, err
		}
		_, err = p.databaseConn.AddCardToDeck(card_id, deck_id)
		if err != nil {
			fmt.Println("Error: while inserting card into deck", err)
			return false, err
		}
	}

	for _, card := range deck.Extradeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		deck_id, err := strconv.Atoi(p.deck_id)
		if err != nil {
			fmt.Println("Error: while converting deck ID in extraDeck to int", err)
			return false, err
		}
		_, err = p.databaseConn.AddCardToDeck(card_id, deck_id)
		if err != nil {
			fmt.Println("Error: while inserting card into deck", err)
			return false, err
		}
	}

	for _, card := range deck.Sidedeck {
		card_id, err := strconv.Atoi(card)
		if err != nil {
			fmt.Println("Error: while converting card ID to int", err)
			return false, err
		}
		deck_id, err := strconv.Atoi(p.deck_id)
		if err != nil {
			fmt.Println("Error: while converting deck ID in sideDeck to int", err)
			return false, err
		}
		_, err = p.databaseConn.AddCardToDeck(card_id, deck_id)
		if err != nil {
			fmt.Println("Error: while inserting card into deck", err)
			return false, err
		}
	}
	fmt.Println("Deck ID:", p.deck_id, "successfully upserted")
	// process all cards in the deck
	return true, nil
}
