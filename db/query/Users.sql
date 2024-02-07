
-- name: GetUserByID :one
SELECT * FROM users
WHERE userid = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;


-- name: UpdateUsers :exec
UPDATE Users
set username = $2,
    password = $3,
    email= $4,
    usertypefk=$5
WHERE userid = $1;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE userid = $1;