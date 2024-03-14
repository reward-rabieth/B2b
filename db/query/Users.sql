-- name: CreateUser :one
INSERT INTO users (
        user_id,username,password,email,role_id
) VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;



-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;


-- name: UpdateUsers :exec
UPDATE users
set username = $2,
    password = $3,
    email= $4,
    role_id=$5
WHERE user_id = $1;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE user_id = $1;

