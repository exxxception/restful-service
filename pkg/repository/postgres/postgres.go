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

func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}
