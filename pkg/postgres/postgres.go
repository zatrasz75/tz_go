package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"zatrasz75/tz_go/pkg/logger"

	"io/ioutil"
	"log"
	"strings"
	"time"
)

// Postgres Хранилище данных
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Pool *pgxpool.Pool
}

func New(connStr string, l logger.LoggersInterface, opts ...Option) (*Postgres, error) {
	pg := &Postgres{}

	// Пользовательские параметры
	for _, opt := range opts {
		opt(pg)
	}

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		l.Info("Postgres пытается подключиться, попыток осталось: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// Close Закрыть
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

// Migrate Миграция таблиц
func (p *Postgres) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	migrationScript, err := ioutil.ReadFile("initScriptPostgres/up.sql")
	if err != nil {
		log.Fatal(err)
	}
	migrationScriptStr := string(migrationScript)

	statements := strings.Split(migrationScriptStr, ";")

	for _, statement := range statements {
		if strings.TrimSpace(statement) != "" {
			_, err = p.Pool.Exec(ctx, statement)
			if err != nil {
				return fmt.Errorf("не удалось прочитать сценарий миграции: %w", err)
			}
		}
	}
	return nil
}
