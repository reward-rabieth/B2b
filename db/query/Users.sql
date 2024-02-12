-- name: CreateUser :one
INSERT INTO users (
        userid,username,password,email,usertypefk
) VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE userid = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;



-- name: GetUserTypeByID :one
SELECT "UserType" FROM "UserTypes" WHERE "UserTypePK" = $1;

-- name: GetUserTypes :many
SELECT "UserTypePK", "UserType" FROM "UserTypes";

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;


-- name: UpdateUsers :exec
UPDATE users
set username = $2,
    password = $3,
    email= $4,
    usertypefk=$5
WHERE userid = $1;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE userid = $1;

