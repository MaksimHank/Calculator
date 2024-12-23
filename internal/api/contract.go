package api

import "calculator/internal/model"

type CalculatorService interface {
	Add(model.Operands) model.Result
	Subtraction(model.Operands) model.Result
	Multiplication(model.Operands) model.Result
	Division(model.Operands) model.Result
}
