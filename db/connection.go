package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var PostgressConnection *PostgresStore

type PostgresStore struct {
	databaseUrl string
	Db          *sqlx.DB
}

func InitPostgresStore(databaseUrl string) {
	PostgressConnection = &PostgresStore{databaseUrl: databaseUrl}
}

func (s *PostgresStore) Connect() error {
	db, err := sqlx.Connect("postgres", s.databaseUrl)

	if err != nil {
		return err
	}

	s.Db = db

	return nil
}

func (s *PostgresStore) Close() error {
	return s.Db.Close()
}
