package main

import (
	"encoding/json"
	"fmt"
	utils "go-rec-sys/libs/utils/class"
	database "go-rec-sys/libs/utils/database"
	"io"
	"net/http"
	"strings"
)

// https://ygoprodeck.com/api/tournament/getTournament.php?id=1968/

type processTournament struct {
	tournamentID string
	databaseConn *database.DatabaseManager
}
type ProcessTournament interface {
	getURL() string
	GetDeck() []utils.Deck
	// NewProcessTournament() *processTournament
	extractDeckURLs(html string) []string
	upsertDeck() []int
}

func NewProcessTournament(id string) *processTournament {
	databaseConn, err := database.NewDatabaseManager("go_rec_sys")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &processTournament{
		tournamentID: id,
		databaseConn: databaseConn,
	}
}

func (pt *processTournament) upsertDeck() []int {
	// Send GET request to the tournament URL
	res, err := http.Get("https://ygoprodeck.com/api/tournament/getTournament.php?id=" + pt.tournamentID)
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
	var tournament utils.Tournament
	// TODO: accept tournament data, with one column to verify if the tournament is already processed
	err = json.Unmarshal(body, &tournament)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	fmt.Println("khiem 1")
	// for every deck URL, create a new processDeck object -> do upsert operation -> return list of deckID in RDS
	for _, deck := range tournament.Listings {
		if deck.PrettyURL != nil {
			// Create a new processDeck object
			url_splited := strings.Split(*deck.PrettyURL, "-")
			deck_id := url_splited[len(url_splited)-1]
			fmt.Println("deck_id:", deck_id)
			// upsert operation, can ultilize go routine to speed up the process by running concurrently all the upsert operation
			pd := NewProcessDeck(deck_id)
			pd.upsert()
		}

	}
	return nil
}

// func main() {
// 	// https://ygoprodeck.com/tournament/niagara-falls-wcq-regional-1935
// 	pt := NewProcessTournament()
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

// }
