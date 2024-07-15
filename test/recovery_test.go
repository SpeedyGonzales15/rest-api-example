package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"rest-api-example/internal/config"
	"rest-api-example/pkg/middleware"
	"testing"
)

func TestRecovery(t *testing.T) {
	config, err := config.LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	middleware.Recovery(mux, config.Server.LogFilePath).ServeHTTP(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("want %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
