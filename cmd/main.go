package main

import (
	"fmt"
	"os"
	"path/filepath"
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/internal/app"
	"zatrasz75/tz_go/pkg/logger"

	_ "zatrasz75/tz_go/docs"
)

func main() {
	l := logger.NewLogger()

	// Получаем текущий рабочий каталог
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущего рабочего каталога:", err)
		return
	}
	// Построение абсолютного пути к файлу configs.yml
	configPath := filepath.Join(cwd, "configs", "configs.yml")

	// Configuration
	cfg, err := configs.NewConfig(configPath)
	if err != nil {
		l.Fatal("ошибка при разборе конфигурационного файла", err)
	}
	app.Run(cfg, l)
}
