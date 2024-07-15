package middleware

import (
	"log"
	"net/http"
	"os"
)

func Recovery(next http.Handler, logFilePAth string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logFilePath := logFilePAth
				file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					log.Fatalf("Ошибка открытия файла: %v", err)
				}
				defer file.Close()

				log.SetOutput(file)
				log.Printf("Паника: %v", err)
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
