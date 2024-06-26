// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	UserID    int32          `json:"user_id"`
	OwnerName string         `json:"owner_name"`
	Balance   int64          `json:"balance"`
	CreatedAt sql.NullTime   `json:"created_at"`
	Currency  sql.NullString `json:"currency"`
}

type Entry struct {
	ID        int64        `json:"id"`
	AccountID int32        `json:"account_id"`
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Transfer struct {
	ID          int64        `json:"id"`
	FromAccount int32        `json:"from_account"`
	ToAccount   int32        `json:"to_account"`
	Amount      int64        `json:"amount"`
	CreatedAt   sql.NullTime `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
