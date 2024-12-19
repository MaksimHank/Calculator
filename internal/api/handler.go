package api

import (
	"calculator/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

type Handler struct {
	calcService CalculatorService
}

func New(calc CalculatorService) *Handler {
	return &Handler{
		calcService: calc,
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем данные из Body
	jsn, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Конвертируем данные []byte -> struct
	var data model.Operands
	err = json.Unmarshal(jsn, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result := h.calcService.Add(data)

	dataBytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// TODO Когда будем созранять данные поменять статус на 201
	w.WriteHeader(http.StatusOK)
	w.Write(dataBytes)
}

func (h *Handler) Subtraction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем данные из Body
	jsn, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Конвертируем данные []byte -> struct
	var data model.Operands
	err = json.Unmarshal(jsn, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result := h.calcService.Subtraction(data)

	dataBytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	// TODO Когда будем созранять данные поменять статус на 201
	w.WriteHeader(http.StatusOK)
	w.Write(dataBytes)
}
