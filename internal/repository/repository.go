package repository

import (
	"context"
	"database/sql"
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

// DeleteCarsById Удаление по id
func (s *Store) DeleteCarsById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// Начать транзакцию
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	query := "SELECT owner_id FROM cars WHERE id = $1"
	var ownerId int

	row := tx.QueryRow(ctx, query, id)
	err = row.Scan(&ownerId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если строки пустые
			return fmt.Errorf("ни одна строка не была возвращена")
		} else {
			return fmt.Errorf("не удалось получить owner по идентификатору: %w", err)
		}
	}

	// Удаляем автомобиль по его идентификатору
	deleteCar := "DELETE FROM cars WHERE id = $1"
	_, err = tx.Exec(ctx, deleteCar, id)
	if err != nil {
		return fmt.Errorf("не удалось автомобиль по идентификатору: %w", err)
	}

	// Удаляем владельца автомобиля
	deleteOwner := "DELETE FROM people WHERE id = $1"
	_, err = tx.Exec(ctx, deleteOwner, ownerId)
	if err != nil {
		return fmt.Errorf("не удалось удалить владельца автомобиля: %w", err)
	}

	// Фиксация транзакции
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
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
