package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "go-rec-sys/db/sqlc"
	utils "go-rec-sys/libs/utils/class"
	database "go-rec-sys/libs/utils/database"
	"io"
	"net/http"
	"strings"
	"time"
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

func NewProcessTournament(id string, databaseConn *database.DatabaseManager) *processTournament {
	return &processTournament{
		tournamentID: id,
		databaseConn: databaseConn,
	}
}

func (pt *processTournament) processTournament() bool {
	// Send GET request to the tournament URL
	res, err := http.Get("https://ygoprodeck.com/api/tournament/getTournament.php?id=" + pt.tournamentID)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Error:", readErr)
		return false
	}
	var tournament utils.Tournament
	// TODO: accept tournament data, with one column to verify if the tournament is already processed
	err = json.Unmarshal(body, &tournament)
	pt.saveTournament(tournament)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return false
	}
	// for every deck URL, create a new processDeck object -> do upsert operation -> return list of deckID in RDS
	for _, deck := range tournament.Listings {
		if deck.PrettyURL != nil {
			// Create a new processDeck object
			url_splited := strings.Split(*deck.PrettyURL, "-")
			deck_id := url_splited[len(url_splited)-1]
			fmt.Println("deck_id:", deck_id)
			// upsert operation, can ultilize go routine to speed up the process by running concurrently all the upsert operation
			pd := NewProcessDeck(deck_id, pt.tournamentID, pt.databaseConn)
			cur_deck, err := pd.getDeck()
			if err != nil {
				fmt.Println("Error:", err)
				return false
			}
			// upsert deck into the database
			pd.upsertDeck(*cur_deck)
			fmt.Println("Insert deck ", deck_id, "successfully")
			pd.upsertCard(*cur_deck)
			fmt.Println("Insert card into", deck_id, "successfully")
		}

	}
	return true
}

func (pt *processTournament) saveTournament(tournament utils.Tournament) (bool, error) {
	// upsert tournament into the database
	eventDate, err := time.Parse("2006-01-02", tournament.EventDate)
	if err != nil {
		fmt.Println("Error parsing event date:", err)
		return false, err
	}
	raw_tournament_info, err := json.Marshal(tournament)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	data := db.CreateTournamentParams{
		ID:                int64(tournament.ID),
		TournamentName:    tournament.Name,
		Tier:              int32(tournament.Tier),
		EventDate:         eventDate,
		PlayerCount:       sql.NullInt32{Int32: int32(tournament.PlayerCount), Valid: true},
		Format:            tournament.Format,
		RawTournamentInfo: string(raw_tournament_info),
	}
	res, err := pt.databaseConn.AddTournament(data)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	return res, nil
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
