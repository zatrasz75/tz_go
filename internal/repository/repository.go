package repository

import (
	"context"
	"fmt"
	"time"
	"zatrasz75/tz_go/models"
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

// SaveNewCar Сохраняет данные о авто и владельце
func (s *Store) SaveNewCar(car models.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// Начать транзакцию
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	// Вставляем владельца в таблицу people
	ownerInsertQuery := `INSERT INTO people (name, surname, patronymic) VALUES ($1, $2, $3) RETURNING id`
	var ownerID int
	err = tx.QueryRow(ctx, ownerInsertQuery, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic).Scan(&ownerID)
	if err != nil {
		return fmt.Errorf("не удалось вставить владельца: %w", err)
	}

	// Вставьте автомобиль в таблицу cars
	carInsertQuery := `INSERT INTO cars (regNum, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(ctx, carInsertQuery, car.RegNum, car.Mark, car.Model, car.Year, ownerID)
	if err != nil {
		return fmt.Errorf("не удалось вставить автомобиль: %w", err)
	}

	// Фиксация транзакции
	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
}
