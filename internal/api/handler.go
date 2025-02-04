package api

import (
	"calculator/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

func (h *Handler) Multiplication(w http.ResponseWriter, r *http.Request) {
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

	result := h.calcService.Multiplication(data)

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

func (h *Handler) Division(w http.ResponseWriter, r *http.Request) {
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

	result := h.calcService.Division(data)

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

func (h *Handler) Result(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	limitS := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res := h.calcService.Results(limit)

	dataBytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(dataBytes)
}
