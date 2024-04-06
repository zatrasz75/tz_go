package controller

import (
	"github.com/gorilla/mux"
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/internal/repository"
	"zatrasz75/tz_go/pkg/logger"
)

// NewRouter -.
func NewRouter(cfg *configs.Config, l logger.LoggersInterface, repo *repository.Store) *mux.Router {
	r := mux.NewRouter()
	newEndpoint(r, cfg, l, repo)
	return r
}
