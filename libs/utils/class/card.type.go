package utils

type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

type CardImage struct {
	ID              int    `json:"id"`
	ImageURL        string `json:"image_url"`
	ImageURLSmall   string `json:"image_url_small"`
	ImageURLCropped string `json:"image_url_cropped"`
}

type CardPrice struct {
	CardmarketPrice    string `json:"cardmarket_price"`
	TcgplayerPrice     string `json:"tcgplayer_price"`
	EbayPrice          string `json:"ebay_price"`
	AmazonPrice        string `json:"amazon_price"`
	CoolstuffincePrice string `json:"coolstuffinc_price"`
}

type RawCardInfo struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	FrameType     string      `json:"frameType"`
	Desc          string      `json:"desc"`
	Atk           int         `json:"atk,omitempty"`
	Def           int         `json:"def,omitempty"`
	Level         int         `json:"level,omitempty"`
	Race          string      `json:"race"`
	Attribute     string      `json:"attribute,omitempty"`
	Archetype     string      `json:"archetype"`
	YgoprodeckURL string      `json:"ygoprodeck_url"`
	CardSets      []CardSet   `json:"card_sets"`
	CardImages    []CardImage `json:"card_images"`
	CardPrices    []CardPrice `json:"card_prices"`
}

// card_type: monster, spell, trap
// card_group: normal, effect, ritual, fusion, synchro, xyz, pendulum, link
type NormalizedCard struct {
	ID          int
	Name        string
	Type        string
	FrameType   string
	Archetype   string
	Race        string
	Attack      int
	Defense     int
	Desc        string
	Image       string
	CardSets    []string
	RawCardInfo RawCardInfo
}
