//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package calculator

import "calculator/internal/model"

type CalculatorRepository interface {
	Insert(operands model.Operands, operator string, result float64)
	Select(limit int) []model.Storage
}
