-- name: CreateUser :one
INSERT INTO users (username, hashed_password, full_name, email) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;

-- name: GetUserWithAccounts :many
SELECT 
    u.username,
    u.hashed_password,
    u.full_name,
    u.email,
    u.password_changed_at,
    u.created_at AS user_created_at,
    a.id AS account_id,
    a.balance,
    a.currency,
    a.created_at AS account_created_at
FROM 
    users u
LEFT JOIN 
    accounts a ON u.username = a.owner
WHERE 
    u.username = $1;
