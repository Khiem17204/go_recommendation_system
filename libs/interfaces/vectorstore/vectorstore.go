package vectorstore

import "context"

type VectorStore interface {
	SearchSimilarCards(ctx context.Context, cards []string) ([]string, error)
	SearchSimilarDecks(ctx context.Context, cards []string) ([]string, error)
}
