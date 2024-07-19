package carddb

import (
	utils "go-rec-sys/libs/utils/class"
)

// Turn raw form of card data into a card object

type CardHelper struct {
}

func (ch *CardHelper) TransformHTMLToCard(html string) (*utils.Card, error) {
	// Use utils package to transform HTML string into card object
	return nil, nil
}
