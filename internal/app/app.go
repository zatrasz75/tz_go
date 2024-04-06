package app

import (
	"os"
	"os/signal"
	"syscall"
	"zatrasz75/tz_go/configs"
	"zatrasz75/tz_go/internal/controller"
	"zatrasz75/tz_go/internal/repository"
	"zatrasz75/tz_go/pkg/logger"
	"zatrasz75/tz_go/pkg/postgres"
	"zatrasz75/tz_go/pkg/server"
)

func Run(cfg *configs.Config, l logger.LoggersInterface) {
	pg, err := postgres.New(cfg.DataBase.ConnStr, l, postgres.OptionSet(cfg.DataBase.PoolMax, cfg.DataBase.ConnAttempts, cfg.DataBase.ConnTimeout))
	if err != nil {
		l.Fatal("ошибка запуска - postgres.New:", err)
	}
	defer pg.Close()

	err = pg.Migrate()
	if err != nil {
		l.Fatal("ошибка миграции", err)
	}

	repo := repository.New(pg, l)

	router := controller.NewRouter(cfg, l, repo)

	srv := server.New(router, server.OptionSet(cfg.Server.AddrHost, cfg.Server.AddrPort, cfg.Server.ReadTimeout, cfg.Server.WriteTimeout, cfg.Server.IdleTimeout, cfg.Server.ShutdownTime))

	go func() {
		err = srv.Start()
		if err != nil {
			l.Error("Остановка сервера:", err)
		}
	}()

	l.Info("Запуск сервера на http://" + cfg.Server.AddrHost + ":" + cfg.Server.AddrPort)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("принят сигнал прерывания прерывание %s", s.String())
	case err = <-srv.Notify():
		l.Error("получена ошибка сигнала прерывания сервера", err)
	}

	err = srv.Shutdown()
	if err != nil {
		l.Error("не удалось завершить работу сервера", err)
	}
}
