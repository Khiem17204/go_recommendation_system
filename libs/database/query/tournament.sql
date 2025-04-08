-- name: CreateTournament :one
INSERT INTO tournaments (
    id,
    tournament_name,
    tier,
    player_count,
    event_date,
    format,
    raw_tournament_info
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;

-- name: GetTournament :one
SELECT * FROM tournaments 
WHERE id = $1 LIMIT 1;

-- name: ListTournaments :many
SELECT * FROM tournaments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteTournament :exec
DELETE FROM tournaments
WHERE id = $1;