package main

import (
	"fmt"
)

type processDeck struct {
	url string
}

func NewProcessDeck(url string) *processDeck {
	return &processDeck{
		url: url,
	}
}

// return list of ID of all the cards if success, return nil if failed
func (p *processDeck) getCards() {
	// Implement the logic to get the cards from the deck URL
	fmt.Println("Getting cards from", p.url)
	// Your code here...
}

// return the deck metadata if success, return nil if failed
func (p *processDeck) getMetaData() {
	// Implement the logic to get the metadata from the deck URL
	fmt.Println("Getting metadata from", p.url)
	// Your code here...
}

// upsert the deck data into the database
func (p *processDeck) upsert() {
	// Implement the logic to upsert the deck data into the database
	fmt.Println("Upserting deck data from", p.url)
	// Your code here...
}
