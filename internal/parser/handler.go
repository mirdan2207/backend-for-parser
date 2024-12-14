package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

const parsingServiceURL = "http://127.0.0.1:8000/post_code" // Адрес вашего сервиса парсинга

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/execute", executeCode).Methods("POST")
}

type ParseRequest struct {
	Code       string `json:"code"`
	ReturnType string `json:"return_type"`
}

type ParseResponse struct {
	ReceivedJSON map[string]interface{} `json:"received_json"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	var req ParseRequest

	// Парсим запрос
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Отправляем запрос к Python-сервису
	resp, err := sendToParserService(req)
	if err != nil {
		http.Error(w, "Error communicating with parsing service: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Читаем и обрабатываем ответ от сервиса
	var parseResponse ParseResponse
	if err := json.Unmarshal(resp, &parseResponse); err != nil {
		http.Error(w, "Invalid response from parsing service", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(parseResponse)
}

func sendToParserService(req ParseRequest) ([]byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Отправка POST-запроса на Python-сервис
	resp, err := http.Post(parsingServiceURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return body, nil
}
