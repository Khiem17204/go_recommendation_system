package main

import (
	"fmt"
	utils "go-rec-sys/libs/utils/class"
)

// https://ygoprodeck.com/api/decks/getDeckInfo.php?deckId=506208
type processDeck struct {
	id string
}

func NewProcessDeck(id string) *processDeck {
	return &processDeck{
		id: id,

// return list of ID of all the cards if success, return nil if failed
func (p *processDeck) getCards() {
	res, err := http.Get("https://ygoprodeck.com/api/decks/getDeckInfo.php?id=" + p.id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error:", readErr)
		return nil
	}
	var deck utils.Deck
	err = json.Unmarshal(body, &deck)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	return &deck, nil
}

// upsert the deck data into the database
func (p *processDeck) upsert() {
	// Implement the logic to upsert the deck data into the database
	fmt.Println("Upserting deck data from", p.url)
	// Your code here...
}
