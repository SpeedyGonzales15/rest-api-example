package main

import (
	"log"
	"net/http"

	"rest-api-example/internal/handler"
	"rest-api-example/internal/repository"
	"rest-api-example/internal/service"
	"rest-api-example/pkg/middleware"
	"time"
)

func main() {
	err := middleware.LogInit("D:\\rest-api-example\\logs\\app.log")
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}

	db, err := repository.NewDB()
	if err != nil {
		middleware.Log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	mux := handlers.InitRoutes()
	recoveredMux := middleware.Recovery(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      recoveredMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	middleware.Log.Info("Server started on :8080")

	err = srv.ListenAndServe()
	if err != http.ErrServerClosed {
		middleware.Log.Errorf("Server start error: %v", err)
	}
}
