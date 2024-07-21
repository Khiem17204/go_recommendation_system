package main

import (
	"fmt"
	utils "go-rec-sys/libs/utils/class"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type processTournament struct {
	baseURL        string
	tournamentName string
	tournamentID   int
}
type ProcessTournament interface {
	getURL() string
	GetDeck() []utils.Deck
	NewProcessTournament() *processTournament
	extractDeckURLs(html string) []string
	upsertDeck() []int
}

func NewProcessTournament() *processTournament {
	return &processTournament{
		baseURL: "https://ygoprodeck.com",
	}
}

func (pt *processTournament) getURL() string {
	return pt.baseURL + "/tournament/" + strings.ToLower(strings.ReplaceAll(pt.tournamentName, " ", "-")) + "-" + strconv.Itoa(pt.tournamentID)
}

func (pt *processTournament) extractDeckURLs(html string) []string {
	// Regular expression to match href attributes containing "/deck/"
	re := regexp.MustCompile(`href="(/deck/[^"]*)"`)

	// Find all matches
	matches := re.FindAllStringSubmatch(html, -1)

	// Extract the URLs from the matches
	var urls []string
	for _, match := range matches {
		if len(match) > 1 {
			urls = append(urls, pt.baseURL+match[1])
		}
	}
	return urls
}

func (pt *processTournament) upsertDeck() []int {
	// Send GET request to the tournament URL
	res, err := http.Get(pt.getURL())
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
	deckURL := pt.extractDeckURLs(string(body))

	// for every deck URL, create a new processDeck object -> do upsert operation -> return list of deckID in RDS
	for _, url := range deckURL {
		// hash deck name+id to get deckID and insert into RDS

		// Create a new processDeck object
		pd := NewProcessDeck(url)
		// upsert operation, can ultilize go routine to speed up the process by running concurrently all the upsert operation

	}
	return nil
}

func main() {
	// https://ygoprodeck.com/tournament/niagara-falls-wcq-regional-1935
	pt := NewProcessTournament()
	pt.tournamentName = "niagara falls wcq regional"
	pt.tournamentID = 1935
	url := pt.getURL()
	res, err := http.Get(url)
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	// Read the response body
	body, readErr := io.ReadAll(res.Body)
	deckURL := pt.extractDeckURLs(string(body))

	fmt.Println(deckURL[0], readErr)

}
