// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( id, user_name, display_name, password )
VALUES ( ?, ?, ?, ? )
RETURNING id, user_name, display_name, password
`

type CreateUserParams struct {
	ID          string
	UserName    string
	DisplayName string
	Password    string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.UserName,
		arg.DisplayName,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.DisplayName,
		&i.Password,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, display_name, user_name FROM users
WHERE user_name = ? LIMIT 1
`

type GetUserRow struct {
	ID          string
	DisplayName string
	UserName    string
}

func (q *Queries) GetUser(ctx context.Context, userName string) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, userName)
	var i GetUserRow
	err := row.Scan(&i.ID, &i.DisplayName, &i.UserName)
	return i, err
}

const getUserForLogin = `-- name: GetUserForLogin :one
SELECT user_name, password FROM users
WHERE user_name = ? LIMIT 1
`

type GetUserForLoginRow struct {
	UserName string
	Password string
}

func (q *Queries) GetUserForLogin(ctx context.Context, userName string) (GetUserForLoginRow, error) {
	row := q.db.QueryRowContext(ctx, getUserForLogin, userName)
	var i GetUserForLoginRow
	err := row.Scan(&i.UserName, &i.Password)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, user_name, display_name, password FROM users
ORDER BY display_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.DisplayName,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
