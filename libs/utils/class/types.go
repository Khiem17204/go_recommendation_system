package utils

// https://ygoprodeck.com/api/tournament/getTournaments.php
type APIResponse struct {
	Data []Tournament `json:"data"`
}

type Deck struct {
	DeckName     string   `json:"deckname"`
	ID           int      `json:"decknumber"`
	Maindeck     []string `json:"maindeck"`
	Extradeck    []string `json:"extradeck"`
	Sidedeck     []string `json:"sidedeck"`
	Rank         string   `json:"tournamentPlacement"`
	TournamentID string   `json:"tournamentslug"`
}

type Tournament struct {
	ID                       int       `json:"id"`
	Name                     string    `json:"name"`
	Country                  string    `json:"country"`
	EventDate                string    `json:"event_date"`
	Winner                   string    `json:"winner"`
	Format                   string    `json:"format"`
	Slug                     string    `json:"slug"`
	PlayerCount              int       `json:"player_count"`
	IsApproximatePlayerCount int       `json:"is_approximate_player_count"`
	Tier                     int       `json:"tier"`
	Listings                 []Listing `json:"listings"`
}

type Listing struct {
	ListingID int     `json:"listing_id"`
	DeckName  *string `json:"deck_name"`
	PrettyURL *string `json:"pretty_url"`
	User      string  `json:"user"`
	Country   string  `json:"country"`
	Placement string  `json:"placement"`
	Arch1     string  `json:"arch_1"`
	Arch1Img  int     `json:"arch_1_img"`
	Arch2     *string `json:"arch_2"`
	Arch2Img  *int    `json:"arch_2_img"`
	Arch3     *string `json:"arch_3"`
	Arch3Img  *int    `json:"arch_3_img"`
	DeckPrice *string `json:"deck_price"`
}
