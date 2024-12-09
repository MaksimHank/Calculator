package model

type Operands struct {
	First  float64 `json:"first"`
	Second float64 `json:"second"`
}

type Result struct {
	Equal float64 `json:"equal"`
}

type Storage struct {
	FirstOperand  float64
	Operator      string
	SecondOperand float64
	Result        float64
}
