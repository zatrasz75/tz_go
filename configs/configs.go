package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"time"
	"zatrasz75/tz_go/pkg/logger"
)

type Config struct {
	Server struct {
		AddrPort     string        `yaml:"port" env:"APP_PORT" env-description:"Server port" env-default:"4141"`
		AddrHost     string        `yaml:"host" env:"APP_IP" env-description:"Server host" env-default:"0.0.0.0"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"READ_TIMEOUT" env-description:"Server ReadTimeout" env-default:"3s"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"WRITE_TIMEOUT" env-description:"Server WriteTimeout" env-default:"3s"`
		IdleTimeout  time.Duration `yaml:"idle-timeout" env:"IDLE_TIMEOUT" env-description:"Server IdleTimeout" env-default:"6s"`
		ShutdownTime time.Duration `yaml:"shutdown-timeout" env:"SHUTDOWN_TIMEOUT" env-description:"Server ShutdownTime" env-default:"10s"`
	} `yaml:"server"`
	DataBase struct {
		ConnStr string `env:"DB_CONNECTION_STRING" env-description:"db string"`

		Host     string `yaml:"host" env:"HOST_DB" env-description:"db host"`
		User     string `yaml:"username" env:"POSTGRES_USER" env-description:"db username"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-description:"db password"`
		Url      string `yaml:"db-url" env:"URL_DB" env-description:"db url"`
		Name     string `yaml:"db-name" env:"POSTGRES_DB" env-description:"db name"`
		Port     string `yaml:"port" env:"PORT_DB" env-description:"db port"`

		PoolMax      int           `yaml:"pool-max" env:"PG_POOL_MAX" env-description:"db PoolMax" env-default:"2"`
		ConnAttempts int           `yaml:"conn-attempts" env:"PG_CONN_ATTEMPTS" env-description:"db ConnAttempts" env-default:"10"`
		ConnTimeout  time.Duration `yaml:"conn-timeout" env:"PG_TIMEOUT" env-description:"db ConnTimeout" env-default:"2s"`
	} `yaml:"database"`
	Api struct {
		Url string `yaml:"url" env:"API_URL" env-description:"api url"`
	} `yaml:"api"`
}

func NewConfig(l logger.LoggersInterface) (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		l.Warn("системе не удается найти указанный файл .env: - %v", err)
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		l.Error("ошибка .env ", err)
		return nil, err
	}
	if err := cleanenv.ReadConfig("./configs/configs.yml", &cfg); err != nil {
		return nil, err
	}

	cfg.DataBase.ConnStr = initDB(cfg)

	return &cfg, nil
}

func initDB(cfg Config) string {
	if cfg.DataBase.ConnStr != "" {
		return cfg.DataBase.ConnStr
	}
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DataBase.Host,
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.Url,
		cfg.DataBase.Port,
		cfg.DataBase.Name,
	)
}
