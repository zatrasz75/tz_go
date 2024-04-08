package controller

import (
	"github.com/gorilla/mux"
	"zatrasz75/tz_go/configs"
	_ "zatrasz75/tz_go/docs"
	"zatrasz75/tz_go/internal/repository"
	"zatrasz75/tz_go/pkg/logger"
)

// @title Swagger API
// @version 1.0
// @description ТЗ Go - апрель.

// @contact.url https://t.me/Zatrasz
// @contact.email zatrasz@ya.ru

// @host localhost:4141
// @BasePath /

// NewRouter -.
func NewRouter(cfg *configs.Config, l logger.LoggersInterface, repo *repository.Store) *mux.Router {
	r := mux.NewRouter()
	newEndpoint(r, cfg, l, repo)
	return r
}
