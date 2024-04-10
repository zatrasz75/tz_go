package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"zatrasz75/tz_go/pkg/logger"

	"log"
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
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			// Проверяем, что подключение действительно было установлено
			err = pg.Pool.Ping(context.Background())
			if err == nil {
				// Подключение успешно, выходим из цикла
				break
			}
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
func (p *Postgres) Migrate(l logger.LoggersInterface) error {
	// Прочитать миграции из папки:
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	// Преобразование pgxpool.Pool в *sql.DB
	db := stdlib.OpenDBFromPool(p.Pool)

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	l.Info("Применена %d миграция!\n", n)

	return nil
}
