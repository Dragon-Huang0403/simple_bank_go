// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (name, balance, currency)
  VALUES ($1, $2, $3)
RETURNING
  id, name, balance, currency, created_at
`

type CreateAccountParams struct {
	Name     string `json:"name"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Name, arg.Balance, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT
  id, name, balance, currency, created_at
FROM
  accounts
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getListAccounts = `-- name: GetListAccounts :many
SELECT
  id, name, balance, currency, created_at
FROM
  accounts
ORDER BY
  id
LIMIT $1 OFFSET $2
`

type GetListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListAccounts(ctx context.Context, arg GetListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getListAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
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

const updateAccount = `-- name: UpdateAccount :one
UPDATE
  accounts
SET
  balance = $1
WHERE
  id = $2
RETURNING
  id, name, balance, currency, created_at
`

type UpdateAccountParams struct {
	Balance int64 `json:"balance"`
	ID      int64 `json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.Balance, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
