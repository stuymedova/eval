package handler

import (
	"encoding/json"
	"net/http"

	"github.com/stuymedova/eval/pkg/eval"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}
	result, err := eval.Calc(req.Expression)
	if err != nil {
		sendError(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}
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
