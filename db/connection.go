package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	databaseUrl string
	Db          *sqlx.DB
}

func NewPostgresStore(databaseUrl string) *PostgresStore {
	return &PostgresStore{databaseUrl: databaseUrl}
}

func (s *PostgresStore) Connect() error {
	db, err := sqlx.Connect("postgres", s.databaseUrl)

	if err != nil {
		return err
	}

	s.Db = db

	return nil
}

func (s *PostgresStore) Close() {
	s.Db.Close()
}
