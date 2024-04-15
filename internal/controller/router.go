package controller

import (
	"github.com/gorilla/mux"
	"zatrasz75/tz_go/configs"
	_ "zatrasz75/tz_go/docs"
	"zatrasz75/tz_go/internal/repository"
	"zatrasz75/tz_go/pkg/logger"
)

// @title Swagger API tz_go
// @version 1.0
// @description ТЗ Go - апрель.
// @description https://docs.google.com/document/u/0/d/1c0GEgi0svIsg14aNAfpTgPv9te9tcGoRmE4kngyD0ow/mobilebasic

// @contact.name Михаил Токмачев
// @contact.url https://t.me/Zatrasz
// @contact.email zatrasz@ya.ru

// @BasePath /

// NewRouter -.
func NewRouter(cfg *configs.Config, l logger.LoggersInterface, repo *repository.Store) *mux.Router {
	r := mux.NewRouter()
	newEndpoint(r, cfg, l, repo)
	return r
}
