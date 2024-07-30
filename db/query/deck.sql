-- name: CreateDeck :one
INSERT INTO decks (
    id,
    deck_name,
    rank,
    tournament_id,
    raw_deck_info
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetDeck :one
SELECT * FROM decks 
WHERE id = $1 LIMIT 1;

-- name: ListDecks :many
SELECT * FROM decks
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: DeleteDeck :exec
DELETE FROM decks
WHERE id = $1;
