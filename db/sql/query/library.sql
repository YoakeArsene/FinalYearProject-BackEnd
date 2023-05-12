-- name: CreateLibrary :one
INSERT INTO libraries(
    id, user_id, game_id
)
VALUES (
           $1,  $2,  $3
       )
    RETURNING *;

-- name: GetUserLibrary :many
SELECT games.*
FROM libraries
         JOIN games ON libraries.game_id = games.id
WHERE libraries.user_id = $1;

-- name: CheckGameInLibrary :one
SELECT COUNT(*) AS game_count
FROM libraries
WHERE user_id = $1
  AND game_id = $2;


-- name: DeleteLibrary :one
DELETE FROM libraries
WHERE user_id = $1 AND game_id = $2
    RETURNING *;
