package storage

import "zatrasz75/tz_go/models"

type RepositoryInterface interface {
	// SaveNewCar Сохраняет данные о авто и владельце
	SaveNewCar(car models.Car) error
	// DeleteCarsById Удаление по id
	DeleteCarsById(id int) error
}
