-- name: AddCardToDeck :one
INSERT INTO cards_in_deck (
    id,
    card_id,
    deck_id,
    card_count
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetCardsFromDeck :many
SELECT * FROM cards_in_deck
WHERE deck_id = $1;

-- name: GetDecksFromCard :many
SELECT * FROM cards_in_deck
WHERE card_id = $1;

-- name: CountCardInDeck :one
SELECT card_count FROM cards_in_deck
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
