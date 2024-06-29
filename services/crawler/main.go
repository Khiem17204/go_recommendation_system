package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// API for get tournaments data from ygoprodeck.com
// https://ygoprodeck.com/api/tournament/getTournaments.php

type Tournament struct {
	ID                       int    `json:"id"`
	Name                     string `json:"name"`
	Country                  string `json:"country"`
	EventDate                string `json:"event_date"`
	Winner                   string `json:"winner"`
	Format                   string `json:"format"`
	Slug                     string `json:"slug"`
	PlayerCount              int    `json:"player_count"`
	IsApproximatePlayerCount int    `json:"is_approximate_player_count"`
}

type APIResponse struct {
	Data []Tournament `json:"data"`
}

// function to fetch tournament
func fetchTournament() ([]Tournament, error) {
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
	data := APIResponse{}
	// Unmarshal the JSON data into the data variable
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		return nil, fmt.Errorf("error unmarshalling JSON data: %v", jsonErr)
	}
	return data.Data, nil
}

func main() {
	tournament, err := fetchTournament()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Tournaments size", len(tournament))
	fmt.Println(tournament[4].Name)
	fmt.Println(tournament[4].ID)
	// for _, t := range tournament {
	// 	fmt.Println(t.Name)
	// }
}
