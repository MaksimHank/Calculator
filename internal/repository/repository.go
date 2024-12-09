package repository

import (
	"calculator/internal/model"
	"log"
)

type Repository struct {
	storage map[int]model.Storage
}

func New() *Repository {
	return &Repository{
		storage: make(map[int]model.Storage),
	}
}

func (r *Repository) Insert(operands model.Operands, operator string, result float64) {
	index := len(r.storage)
	r.storage[index] = model.Storage{
		FirstOperand:  operands.First,
		Operator:      operator,
		SecondOperand: operands.Second,
		Result:        result,
	}
	log.Println(r.storage)
}
