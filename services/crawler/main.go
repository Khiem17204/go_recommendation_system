package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
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
		fmt.Printf("Error connecting to database: %v\n", err)
		return false
	}
	defer databaseConn.Close()

	for i := 0; i < len(tournaments); i += 1 {
		id := fmt.Sprintf("%d", tournaments[i].ID)
		pt := NewProcessTournament(id, databaseConn)
		pt.processTournament()
	}
	return true
}

func crawlDeck() error {
	tournaments, err := fetchAllTournament()
	if err != nil {
		return fmt.Errorf("error fetching tournaments: %v", err)
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
			if !processByBatch(tournaments[start:end]) {
				fmt.Printf("Error processing batch %d to %d\n", start, end-1)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All batches processed")
	return nil
}

func crawlCard() error {
	databaseConn, err := database.NewDatabaseManager("go_rec_sys")
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	defer databaseConn.Close()

	ch := NewCardHelper("https://db.ygoprodeck.com/api/v7/cardinfo.php")
	fmt.Println("Starting card crawling...")
	success, err := ch.saveAllCards()
	if err != nil {
		return fmt.Errorf("error saving cards: %v", err)
	}
	if !success {
		return fmt.Errorf("failed to save all cards")
	}
	fmt.Println("Finished card crawling")
	return nil
}

func main() {
	// Parse command line arguments
	mode := flag.String("mode", "", "Crawler mode: 'card' or 'deck'")
	flag.Parse()

	if *mode == "" {
		fmt.Println("Please specify a mode: --mode=card or --mode=deck")
		os.Exit(1)
	}

	// Run the appropriate crawler based on mode
	var err error
	switch *mode {
	case "card":
		err = crawlCard()
	case "deck":
		err = crawlDeck()
	default:
		fmt.Printf("Invalid mode: %s\n", *mode)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error during crawling: %v\n", err)
		os.Exit(1)
	}
}
