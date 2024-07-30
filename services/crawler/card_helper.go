package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "go-rec-sys/db/sqlc"
	utils "go-rec-sys/libs/utils/class"
	database "go-rec-sys/libs/utils/database"
	"net/http"
)

type cardHelper struct {
	url          string
	databaseConn *database.DatabaseManager
}

type CardHelper interface {
	saveAllCards() (bool, error)
}

func (ch *cardHelper) getCardsInfo() ([]utils.NormalizedCard, error) {
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
		normalizedCard := ch.convertToNormalizedCard(rawCard)
		normalizedCards[i] = normalizedCard
	}

	return normalizedCards, nil
}

func (ch *cardHelper) convertToNormalizedCard(rawCard utils.RawCardInfo) utils.NormalizedCard {
	attack := -1
	defense := -1
	level := -1
	if rawCard.Atk != nil {
		attack = *rawCard.Atk
	}
	if rawCard.Def != nil {
		defense = *rawCard.Def
	}
	if rawCard.Level != nil {
		level = *rawCard.Level
	}
	normalizedCard := utils.NormalizedCard{
		ID:          rawCard.ID,
		Name:        rawCard.Name,
		Type:        rawCard.Type,
		Archetype:   rawCard.Archetype,
		Race:        rawCard.Race,
		Attribute:   rawCard.Attribute,
		Desc:        rawCard.Desc,
		Level:       level,
		RawCardInfo: rawCard,
		Attack:      attack,
		Defense:     defense,
	}
	return normalizedCard
}

func (ch *cardHelper) saveCard(card utils.NormalizedCard) (bool, error) {
	// Save the card to the database
	archetype := sql.NullString{String: card.Archetype, Valid: card.Archetype != ""}
	attribute := sql.NullString{String: card.Attribute, Valid: card.Attribute != ""}
	race := sql.NullString{String: card.Race, Valid: card.Race != ""}
	attack := sql.NullInt32{Int32: int32(card.Attack), Valid: card.Attack != -1}
	defense := sql.NullInt32{Int32: int32(card.Defense), Valid: card.Defense != -1}
	level := sql.NullInt32{Int32: int32(card.Level), Valid: card.Level != -1}
	rawCardInfoJSON, err := json.Marshal(card.RawCardInfo)
	if err != nil {
		return false, fmt.Errorf("failed to marshal rawCardInfo: %v", err)
	}
	data := db.CreateCardParams{
		ID:          int64(card.ID),
		Name:        card.Name,
		Type:        card.Type,
		FrameType:   card.RawCardInfo.FrameType,
		Archetype:   archetype,
		Attribute:   attribute,
		Race:        race,
		Level:       level,
		Attack:      attack,
		Defense:     defense,
		Description: card.Desc,
		RawCardInfo: rawCardInfoJSON,
	}

	res, err := ch.databaseConn.AddCard(data)
	if err != nil {
		return false, err
	}
	fmt.Println("Insert card ", card.ID, "successfully")
	return res, nil
}

func (ch *cardHelper) saveAllCards() (bool, error) {
	cards, err := ch.getCardsInfo()
	if err != nil {
		return false, err
	}
	defer ch.databaseConn.Close()
	for _, card := range cards {
		_, err := ch.saveCard(card)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func NewCardHelper(url string) CardHelper {
	databaseConn, err := database.NewDatabaseManager("go_rec_sys")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}
	return &cardHelper{
		url:          url,
		databaseConn: databaseConn,
	}
}
