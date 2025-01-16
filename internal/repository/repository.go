package repository

import (
	"calculator/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	storage map[int]model.Storage
	conn    *sqlx.DB
}

func New() *Repository {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 4444, "master", "master", "master")

	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		storage: make(map[int]model.Storage),
		conn:    conn,
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

	query := "INSERT INTO results (first_operand, operator, second_operand, result) VALUES ($1, $2, $3, $4)"

	_, err := r.conn.Exec(query, operands.First, operator, operands.Second, result)
	if err != nil {
		log.Println("failed to insert into results table")
	}
}

func (r *Repository) Select(limit int) []model.Storage {
	query := "SELECT first_operand, operator, second_operand, result FROM results LIMIT $1"

	var results []model.Storage
	err := r.conn.Select(&results, query, limit)
	if err != nil {
		log.Println("failed to select from results table")
	}
	return results
}
