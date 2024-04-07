package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
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

// UpdateCarsById Изменение одного или нескольких полей по идентификатору
func (s *Store) UpdateCarsById(car models.Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// Начать транзакцию
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	// Получаем внешний ключ для владельца cars
	query := "SELECT owner_id FROM cars WHERE id = $1"
	var ownerId int

	row := tx.QueryRow(ctx, query, car.ID)
	err = row.Scan(&ownerId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если строки пустые
			return fmt.Errorf("ни одна строка не была возвращена")
		} else {
			return fmt.Errorf("не удалось получить owner по идентификатору: %w", err)
		}
	}

	// Динамически создаём SQL-запрос для обновления данных cars
	var queryCar strings.Builder
	queryCar.Reset()
	queryCar.WriteString("UPDATE cars SET ")
	var args []interface{}
	var argIndex int
	var fieldsUpdated bool

	if car.RegNum != "" {
		queryCar.WriteString("regNum = $" + strconv.Itoa(argIndex+1))
		args = append(args, car.RegNum)
		argIndex++
		fieldsUpdated = true
	}
	if car.Mark != "" {
		if fieldsUpdated {
			queryCar.WriteString(", ")
		}
		queryCar.WriteString("mark = $" + strconv.Itoa(argIndex+1))
		args = append(args, car.Mark)
		argIndex++
		fieldsUpdated = true
	}
	if car.Model != "" {
		if fieldsUpdated {
			queryCar.WriteString(", ")
		}
		queryCar.WriteString("model = $" + strconv.Itoa(argIndex+1))
		args = append(args, car.Model)
		argIndex++
		fieldsUpdated = true
	}
	if car.Year != 0 {
		if fieldsUpdated {
			queryCar.WriteString(", ")
		}
		queryCar.WriteString("year = $" + strconv.Itoa(argIndex+1))
		args = append(args, car.Year)
		argIndex++
		fieldsUpdated = true
	}

	// Добавляем WHERE только в том случае, если какие-либо поля были обновлены
	if fieldsUpdated {
		queryCar.WriteString(" WHERE id = $" + strconv.Itoa(argIndex+1))
		args = append(args, car.ID)

		_, err = tx.Exec(ctx, queryCar.String(), args...)
		if err != nil {
			s.l.Error("не удалось обновить автомобиль:", err)
			return fmt.Errorf("не удалось обновить автомобиль: %w", err)
		}
	}

	// Динамически создаём SQL-запрос для обновления данных people
	var queryOwner strings.Builder
	queryOwner.Reset()
	queryOwner.WriteString("UPDATE people SET ")
	var ownerArgs []interface{}
	var ownerArgIndex int
	var ownerFieldsUpdated bool

	if car.Owner.Name != "" {
		queryOwner.WriteString("name = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, car.Owner.Name)
		ownerArgIndex++
		ownerFieldsUpdated = true
	}
	if car.Owner.Surname != "" {
		if ownerFieldsUpdated {
			queryOwner.WriteString(", ")
		}
		queryOwner.WriteString("surname = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, car.Owner.Surname)
		ownerArgIndex++
		ownerFieldsUpdated = true
	}
	if car.Owner.Patronymic != "" {
		if ownerFieldsUpdated {
			queryOwner.WriteString(", ")
		}
		queryOwner.WriteString("patronymic = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, car.Owner.Patronymic)
		ownerArgIndex++
		ownerFieldsUpdated = true
	}

	// Добавляем WHERE только в том случае, если какие-либо поля были обновлены
	if ownerFieldsUpdated {
		queryOwner.WriteString(" WHERE id = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, ownerId)

		_, err = tx.Exec(ctx, queryOwner.String(), ownerArgs...)
		if err != nil {
			s.l.Error("не удалось обновить владельца:", err)
			return fmt.Errorf("не удалось обновить владельца: %w", err)
		}
	}

	// Фиксация транзакции
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
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
