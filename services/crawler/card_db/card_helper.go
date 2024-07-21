package main

import (
	"encoding/json"
	"fmt"
	utils "go-rec-sys/libs/utils/class"
	"net/http"
	"strconv"
)

type cardHelper struct {
	url string
}

type CardHelper interface {
	getCardInfo() ([]utils.NormalizedCard, error)
}

func (ch *cardHelper) getCardInfo() ([]utils.NormalizedCard, error) {
	resp, err := http.Get(ch.url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	var rawResponse struct {
		Data []utils.RawCardInfo `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&rawResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	// Convert rawCardInfo to utils.NormalizedCard
	normalizedCards := make([]utils.NormalizedCard, len(rawResponse.Data))
	for i, rawCard := range rawResponse.Data {
		normalizedCard := convertToNormalizedCard(rawCard)
		normalizedCards[i] = normalizedCard
	}

	return normalizedCards, nil
}

func convertToNormalizedCard(rawCard utils.RawCardInfo) utils.NormalizedCard {
	normalizedCard := utils.NormalizedCard{
		ID:          rawCard.ID,
		Name:        rawCard.Name,
		Type:        rawCard.Type,
		Archetype:   rawCard.Archetype,
		Race:        rawCard.Race,
		RawCardInfo: rawCard,
	}
	if rawCard.Type == "Monster" {
		normalizedCard.Attack = rawCard.Atk
		normalizedCard.Defense = rawCard.Def
	} else {
		normalizedCard.Attack = -1
		normalizedCard.Defense = -1
	}
	return normalizedCard
}

func main() {
	ch := cardHelper{
		url: "https://db.ygoprodeck.com/api/v7/cardinfo.php",
	}

	cards, err := ch.getCardInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Card ID: %s\n", strconv.Itoa(cards[0].ID))
	fmt.Printf("Card Name: %s\n", cards[0].Name)
	fmt.Printf("Card Type: %s\n", cards[0].Type)
	fmt.Printf("Card Archetype: %s\n", cards[0].Archetype)
	fmt.Printf("Card Race: %s\n", cards[0].Race)
	fmt.Printf("Card Attack: %d\n", cards[0].Attack)
	fmt.Printf("Card Defense: %d\n", cards[0].Defense)
	fmt.Println("--------------------")
}
