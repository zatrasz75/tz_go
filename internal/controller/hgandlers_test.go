package controller

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/models"
)

// Mock server handler
func handler(w http.ResponseWriter, r *http.Request) {
	// Пример response
	car := models.Car{
		RegNum: "X123XX150",
		Mark:   "Lada",
		Model:  "Vesta",
		Year:   2002,
		Owner: models.People{
			Name:       "Василий",
			Surname:    "Васин",
			Patronymic: "Васильевич",
		},
	}

	json.NewEncoder(w).Encode(car)
}

func TestGetAvailableReleases(t *testing.T) {
	// мок-сервер
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// экземпляр структуры api с мок-сервером
	apiInstance := &api{
		Cfg: &configs.Config{
			Api: struct {
				Url string `yaml:"url" env:"EXTERNAL_API_URL" env-description:"api url"`
			}(struct {
				Url string `yaml:"url" env:"API_URL" env-description:"api url"`
			}(struct {
				Url string
			}{
				Url: server.URL, //URL мок-сервера
			})),
		},
	}

	// Вызов функции getCarInfo, передав в неё регистрационный номер автомобиля
	regNum := "X123XX150"
	carInfo, err := apiInstance.carInfo(regNum)
	if err != nil {
		t.Fatalf("Ошибка при получении информации об автомобиле: %v", err)
	}
	t.Log("carInfo", carInfo)
	fmt.Printf("Информация об автомобиле: %+v\n", carInfo)
}
