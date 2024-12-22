package main

import (
	"net/http"
	"os"

	"github.com/stuymedova/eval/internal/handler"
	"github.com/stuymedova/eval/internal/logger"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handler.CalcHandler)

	logger.Logger.Info("server starting", "port", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
