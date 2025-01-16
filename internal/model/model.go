package model

type Operands struct {
	First  float64 `json:"first"`
	Second float64 `json:"second"`
}

type Result struct {
	Equal float64 `json:"equal"`
}

type Storage struct {
	FirstOperand  float64 `db:"first_operand"`
	Operator      string  `db:"operator"`
	SecondOperand float64 `db:"second_operand"`
	Result        float64 `db:"result"`
}

type Results struct {
	List []Storage `json:"list"`
}
