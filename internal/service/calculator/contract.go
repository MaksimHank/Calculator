package calculator

import "calculator/internal/model"

type CalculatorRepository interface {
	Insert(operands model.Operands, operator string, result float64)
}
