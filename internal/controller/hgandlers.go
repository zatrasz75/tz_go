package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/internal/repository"
	"zatrasz75/tz_go/internal/storage"
	"zatrasz75/tz_go/models"
	"zatrasz75/tz_go/pkg/logger"
)

type api struct {
	Cfg  *configs.Config
	l    logger.LoggersInterface
	repo storage.RepositoryInterface
}

func newEndpoint(r *mux.Router, cfg *configs.Config, l logger.LoggersInterface, repo *repository.Store) {
	en := &api{cfg, l, repo}
	r.HandleFunc("/", en.home).Methods(http.MethodGet)
	r.HandleFunc("/cars", en.addCars).Methods(http.MethodPost)
	r.HandleFunc("/cars/{id}", en.deleteCarsById).Methods(http.MethodDelete)
}

func (a *api) home(w http.ResponseWriter, _ *http.Request) {
	// Устанавливаем правильный Content-Type для HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://frontend.host")

	// Выводим дополнительную строку на страницу
	str := []byte("Добро пожаловать! ")

	_, err := fmt.Fprintf(w, "<p>%s</p>", str)
	if err != nil {
		http.Error(w, "Ошибка записи на страницу", http.StatusInternalServerError)
		a.l.Error("Ошибка записи на страницу", err)
	}
}

func (a *api) deleteCarsById(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		a.l.Debug("Отсутствует идентификатор в запросе")
		http.Error(w, "Отсутствует идентификатор в запросе", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		a.l.Error("не удалось преобразовать строку в число", err)
		http.Error(w, "не удалось преобразовать строку в число", http.StatusBadRequest)
		return
	}

	err = a.repo.DeleteCarsById(id)
	if err != nil {
		a.l.Error("Ошибка при удалении данных", err)
		http.Error(w, "Ошибка при удалении данных", http.StatusInternalServerError)
		return
	}
	a.l.Info("Данные c id %d успешно удалены", id)

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Данные автомобили успешно удалены"))
	if err != nil {
		http.Error(w, "ошибка при отправке данных", http.StatusInternalServerError)
		a.l.Error("ошибка при отправке данных: ", err)
		return
	}
}

func (a *api) addCars(w http.ResponseWriter, r *http.Request) {
	var request struct {
		RegNums []string `json:"regNums"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "не удалось проанализировать запрос JSON", http.StatusBadRequest)
		a.l.Error("не удалось проанализировать запрос JSON", err)
		return
	}

	for _, regNum := range request.RegNums {
		carInfo, err := a.getCarInfo(regNum)
		if err != nil {
			http.Error(w, "недопустимый текст запроса", http.StatusInternalServerError)
			a.l.Error("Не удалось расшифровать текст запроса", err)
			return
		}

		err = a.repo.SaveNewCar(carInfo)
		if err != nil {
			a.l.Error("Ошибка при добавлении данных", err)
			http.Error(w, "Ошибка при добавлении данных", http.StatusInternalServerError)
			return
		}
		a.l.Info("Информация об автомобиле для %s: %v сохранена", regNum, carInfo)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Автомобили успешно добавлены"))
	if err != nil {
		http.Error(w, "ошибка при отправке данных", http.StatusInternalServerError)
		a.l.Error("ошибка при отправке данных: ", err)
		return
	}
}

func (a *api) getCarInfo(regNum string) (models.Car, error) {
	var car models.Car
	car.RegNum = regNum
	url := fmt.Sprintf(a.Cfg.Api.Url + regNum)

	// На случай https , отключаем проверку сертификата
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return car, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return car, fmt.Errorf("не удалось обаготить данные о авто", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return car, err
	}

	err = json.Unmarshal(body, &car)
	if err != nil {
		return car, err
	}

	return car, nil
}
