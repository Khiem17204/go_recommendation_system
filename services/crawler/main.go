package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	utils "github.com/Khiem17204/go_recommendation_system/libs/utils/class"
	database "github.com/Khiem17204/go_recommendation_system/libs/utils/database"
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

func processByBatch(tournaments []utils.Tournament) bool {
	databaseConn, err := database.NewDatabaseManager("go_rec_sys")
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	for i := 0; i < len(tournaments); i += 1 {
		id := fmt.Sprintf("%d", tournaments[i].ID)
		pt := NewProcessTournament(id, databaseConn)
		pt.processTournament()
	}
	databaseConn.Close()
	return true
}

func crawlDeck() {
	tournaments, err := fetchAllTournament()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Process by batch of 20, do parallel processing
	for i := 0; i < len(tournaments); i += 20 {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + 20
			if end > len(tournaments) {
				end = len(tournaments)
			}
			fmt.Printf("Processing batch %d to %d\n", start, end-1)
			processByBatch(tournaments[start:end])
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All batches processed")
}

func crawlCard() {
	ch := NewCardHelper("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	fmt.Println("start upsert card")
	ch.saveAllCards()
	fmt.Println("finish upsert card")
}

func main() {
	crawlCard()
	crawlDeck()
}
