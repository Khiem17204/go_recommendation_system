package main

import (
	"encoding/json"
	"fmt"
	utils "go-rec-sys/libs/utils/class"
	"io"
	"net/http"
)

// API for get tournaments data from ygoprodeck.com
// https://ygoprodeck.com/api/tournament/getTournaments.php

// function to fetch tournament -> move to new file: processTournament.go
func fetchAllTournament() ([]utils.Tournament, error) {
	// Send GET request to the https://ygoprodeck.com/api/tournament/getTournaments.php
	res, err := http.Get("https://ygoprodeck.com/api/tournament/getTournaments.php")

	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, fmt.Errorf("error reading response body: %v", readErr)
	}
	data := utils.APIResponse{}
	// Unmarshal the JSON data into the data variable
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data: %v", jsonErr)
	}
	return data.Data, nil
}

// function to fetch tournament detail: all deck -> move to new file: processTournament.go
func fetchTournament(id string) ([]utils.Deck, error) {
	// construct endpoint
	// get process endpoint
	// parse json into deck object
	// return []deck object
	return nil, nil
}

func main() {
	// https://ygoprodeck.com/tournament/niagara-falls-wcq-regional-1935
	pt := NewProcessTournament("1935")
	fmt.Println("start upsert deck")
	pt.upsertDeck()
	fmt.Println("finish upsert deck")
	// 	pt.tournamentName = "niagara falls wcq regional"
	// 	pt.tournamentID = 1935
	// 	url := pt.getURL()
	// 	res, err := http.Get(url)
	// 	if err != nil {
	// 		// Handle the error
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	defer res.Body.Close()
	// 	// Read the response body
	// 	body, readErr := io.ReadAll(res.Body)
	// 	deckURL := pt.extractDeckURLs(string(body))

	// 	fmt.Println(deckURL[0], readErr)

}
