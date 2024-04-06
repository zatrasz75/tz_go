package repository

import (
	"zatrasz75/tz_go/pkg/logger"
	"zatrasz75/tz_go/pkg/postgres"
)

type Store struct {
	*postgres.Postgres
	l logger.LoggersInterface
}

func New(pg *postgres.Postgres, l logger.LoggersInterface) *Store {
	return &Store{pg, l}
}
