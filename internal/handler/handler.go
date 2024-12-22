package handler

import (
	"encoding/json"
	"net/http"

	"github.com/stuymedova/eval/pkg/eval"
	"github.com/stuymedova/eval/internal/logger"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Info("received request",
		"method", r.Method,
		"remote_addr", r.RemoteAddr,
		"path", r.URL.Path,
	)

	if r.Method != http.MethodPost {
		logger.Logger.Error("method not allowed", "method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Logger.Error("invalid request body", "error", err)
		sendError(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	logger.Logger.Info("processing expression", "expression", req.Expression)
	result, err := eval.Calc(req.Expression)
	if err != nil {
		logger.Logger.Error("expression evaluation failed",
			"expression", req.Expression,
			"error", err,
		)
		sendError(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	logger.Logger.Info("calculation successful",
		"expression", req.Expression,
		"result", result,
	)

	response := CalcResponse{
		Result: result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(CalcResponse{
		Error: message,
	})
}
