package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		body         interface{}
		expectedCode int
		expectedBody CalcResponse
	}{
		{
			name:         "valid expression",
			method:       http.MethodPost,
			body:         CalcRequest{Expression: "2+2*2"},
			expectedCode: http.StatusOK,
			expectedBody: CalcResponse{Result: 6},
		},
		{
			name:         "invalid expression",
			method:       http.MethodPost,
			body:         CalcRequest{Expression: "2++2"},
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: CalcResponse{Error: "Expression is not valid"},
		},
		{
			name:         "empty expression",
			method:       http.MethodPost,
			body:         CalcRequest{Expression: ""},
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: CalcResponse{Error: "Expression is not valid"},
		},
		{
			name:         "invalid method",
			method:       http.MethodGet,
			body:         nil,
			expectedCode: http.StatusMethodNotAllowed,
			expectedBody: CalcResponse{},
		},
		{
			name:         "invalid json",
			method:       http.MethodPost,
			body:         "not a json",
			expectedCode: http.StatusUnprocessableEntity,
			expectedBody: CalcResponse{Error: "Invalid request body"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reqBody []byte
			var err error
			if str, ok := tt.body.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.body)
				if err != nil {
					t.Fatal(err)
				}
			}
			req := httptest.NewRequest(tt.method, "/api/v1/calculate", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			CalcHandler(w, req)
			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}
			if tt.expectedCode == http.StatusMethodNotAllowed {
				return
			}
			var response CalcResponse
			if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
				t.Fatal(err)
			}
			if tt.expectedBody.Error != "" {
				if response.Error != tt.expectedBody.Error {
					t.Errorf("expected error %q, got %q", tt.expectedBody.Error, response.Error)
				}
			} else {
				if response.Result != tt.expectedBody.Result {
					t.Errorf("expected result %f, got %f", tt.expectedBody.Result, response.Result)
				}
			}
		})
	}
}
