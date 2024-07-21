-- name: AddCardToDeck :one
INSERT INTO cards_in_deck (
    card_id,
    deck_id
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetCardsFromDeck :many
SELECT * FROM cards_in_deck
WHERE deck_id = $1;

-- name: GetDecksFromCard :many
SELECT * FROM cards_in_deck
WHERE card_id = $1;

-- name: CountCardInDeck :one
SELECT COUNT(*) FROM cards_in_deck
WHERE card_id = $1 AND deck_id = $2;

-- name: DeleteCardFromDeck :exec
DELETE FROM cards_in_deck
WHERE card_id = $1 AND deck_id = $2;

-- name: DeleteAllCardsFromDeck :exec
DELETE FROM cards_in_deck
WHERE deck_id = $1;

-- name: DeleteAllDecksFromCard :exec
DELETE FROM cards_in_deck
WHERE card_id = $1;
