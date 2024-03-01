package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable      = "users"
	postsTable      = "posts"
	usersPostsTable = "users_posts"
	accountsTable   = "accounts"
)

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	uriCon := "host=localhost port=5432 user=postgres dbname=postgres password=0000 sslmode=disable"

	db, err := sqlx.Open("postgres", uriCon)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}
