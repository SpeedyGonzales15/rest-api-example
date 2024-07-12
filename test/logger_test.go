package test

import (
	"encoding/json"
	"os"
	"rest-api-example/pkg/middleware"
	"testing"
)

func TestLogInit(t *testing.T) {
	logFilePath := "testing.log"
	err := middleware.LogInit(logFilePath)

	if err != nil {
		t.Errorf("Error initializing log: %v", err)

	}
	defer os.Remove(logFilePath)

	middleware.Log.Info("Test message")

	file, err := os.Open(logFilePath)
	if err != nil {
		t.Errorf("Error opening log file: %v", err)
	}
	defer file.Close()

	var logEntry map[string]interface{}
	err = json.NewDecoder(file).Decode(&logEntry)
	if err != nil {
		t.Errorf("Error decoding log entry: %v", err)
	}

	if logEntry["msg"] != "Test message" {
		t.Errorf("Log entry has incorrect message: got %s, want %s", logEntry["msg"], "Test message")
	}
}
