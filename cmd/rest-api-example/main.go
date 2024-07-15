package main

import (
	"log"
	"net/http"

	"rest-api-example/internal/config"
	"rest-api-example/internal/handler"
	"rest-api-example/internal/migration"
	"rest-api-example/internal/repository"
	"rest-api-example/internal/service"
	"rest-api-example/pkg/middleware"
	"time"
)

func main() {
	config, err := config.LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewDB(config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.DbName)
	if err != nil {
		log.Fatal(err)
	}

	err = migration.InitUserTable(db)
	if err != nil {
		log.Fatalf("Ошибка инициализации таблицы: %v", err)
	}
	err = migration.InitProductTable(db)
	if err != nil {
		log.Fatalf("Ошибка инициализации таблицы: %v", err)
	}
	err = migration.InitOrderTable(db)
	if err != nil {
		log.Fatalf("Ошибка инициализации таблицы: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	mux := handlers.InitRoutes()
	loggedMux := middleware.LoggingMiddleware(mux, config.Server.LogFilePath)
	recoveredMux := middleware.Recovery(loggedMux, config.Server.LogFilePath)

	srv := &http.Server{
		Addr:           ":" + config.Server.Port,
		Handler:        recoveredMux,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("Server start error: %v", err)
	}
}
