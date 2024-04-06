package main

import (
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/internal/app"
	"zatrasz75/tz_go/pkg/logger"
)

func main() {
	l := logger.NewLogger()

	// Configuration
	cfg, err := configs.NewConfig(l)
	if err != nil {
		l.Fatal("ошибка при разборе конфигурационного файла", err)
	}
	app.Run(cfg, l)
}
