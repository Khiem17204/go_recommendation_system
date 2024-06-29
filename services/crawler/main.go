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
func fetchTournament(name string, id int) ([]utils.Deck, error) {
	// construct endpoint
	// get process endpoint
	// parse json into deck object
	// return []deck object
	return nil, nil
}

// func main() {
// 	// fetch all tournament
// 	// retrived processed tournament
// 	// process only new tournament
// 	// fetch tournament detail
// 	// process tournament detail
// 	// fetch deck
// 	// process deck -> insert into db

// 	tournament, err := fetchAllTournament()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("Tournaments size", len(tournament))
// 	fmt.Println(tournament[4].Name)
// 	fmt.Println(tournament[4].ID)
// 	// for _, t := range tournament {
// 	// 	fmt.Println(t.Name)
// 	// }
// }
