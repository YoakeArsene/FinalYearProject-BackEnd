-- name: CreateGame :one
INSERT INTO games(
    id,
    name,
    surname,
    price,
    "desc",
    link,
    release,
    platforms,
    genre,
    developers,
    publishers,
    "inCart",
    selected,
    "isHovered",
    "isLiked",
    rating,
    cover,
    footage
) VALUES (
    $1,  $2,  $3,  $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
    CAST($18 AS VARCHAR[])
)
RETURNING *;

-- name: GetAllGames :many
Select * FROM games;

-- name: UpdateGame :one
UPDATE games
SET  name = $1,
     surname = $2,
     price = $3,
     "desc" = $4,
     link = $5,
     release = $6,
     platforms = $7,
     genre = $8,
     developers = $9,
     publishers = $10,
     "inCart" = $11,
     selected = $12,
     "isHovered" = $13,
     "isLiked" = $14,
     rating = $15,
     cover = $16,
     footage = CAST($17 AS VARCHAR[])
WHERE id = $18
    RETURNING *;

-- name: DeleteGame :one
DELETE FROM games
WHERE id = $1
    RETURNING *;

