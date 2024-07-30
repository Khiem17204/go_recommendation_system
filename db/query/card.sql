-- name: CreateCard :one
INSERT INTO cards (
  id,
  name,
  type,
  frame_type,
  archetype,
  attribute,
  race,
  level,
  attack,
  defense,
  description,
  raw_card_info
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12
) RETURNING *;

-- name: GetCard :one
SELECT * FROM cards 
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM cards
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: DeleteCard :exec
DELETE FROM cards
WHERE id = $1;
