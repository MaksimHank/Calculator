package api

import "calculator/internal/model"

type CalculatorService interface {
	Add(model.Operands) model.Result
	Subtraction(model.Operands) model.Result
}
