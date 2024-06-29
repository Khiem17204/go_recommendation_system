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

// APIResponse represents the structure of the API response
type APIResponse struct {
	Data []Tournament `json:"data"`
}

type Deck struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Archetype   string `json:"archetype"`
	Author      string `json:"author"`
	Format      string `json:"format"`
	Slug        string `json:"slug"`
	CardCount   int    `json:"card_count"`
	IsPrivate   bool   `json:"is_private"`
	IsPublished bool   `json:"is_published"`
}
