package utils

// Tournament represents a single tournament's data
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

// https://ygoprodeck.com/api/tournament/getTournaments.php
type APIResponse struct {
	Data []Tournament `json:"data"`
}

type Deck struct {
	url    string
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
