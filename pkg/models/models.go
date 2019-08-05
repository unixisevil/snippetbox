package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int       `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Created time.Time `db:"created"`
	Expires time.Time `db:"expires"`
}

type User struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Email          string `db:"email"`
	HashedPassword []byte
	Created        time.Time `db:"created"`
	Active         bool      `db:"active"`
}
