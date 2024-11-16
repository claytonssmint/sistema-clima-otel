package handlers

import (
	"encoding/json"
	"github.com/claytonssmint/sistema-clima-otel/serviceb/handlers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
)

type CEPRequest struct {
	CEP string `json:"cep"`
}

var tracer trace.Tracer

func init() {
	tracer = otel.Tracer("serviceA")
}

func HandleCEP(w http.ResponseWriter, r *http.Request) {
	var cepReq CEPRequest
	if err := json.NewDecoder(r.Body).Decode(&cepReq); err != nil || len(cepReq.CEP) != 8 {
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	ctx, span := tracer.Start(r.Context(), "ServiceA:ValidateCEP")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://serviceb:8081/weather?cep="+cepReq.CEP, nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write([]byte("can not find zipcode"))
		return
	}

	var response handlers.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
