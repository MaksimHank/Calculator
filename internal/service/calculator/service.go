package calculator

import "calculator/internal/model"

type Service struct {
	calcRepo CalculatorRepository
}

func New(calc CalculatorRepository) *Service {
	return &Service{
		calcRepo: calc,
	}
}

func (s *Service) Add(data model.Operands) model.Result {
	result := data.First + data.Second
	s.calcRepo.Insert(data, "+", result)
	return model.Result{
		Equal: result,
	}
}

func (s *Service) subtraction(data model.Operands) model.Result {
	result := data.First - data.Second
	s.calcRepo.Insert(data, "-", result)
	return model.Result{
		Equal: result,
	}
}

func (s *Service) multiplication(data model.Operands) model.Result {
	result := data.First * data.Second
	s.calcRepo.Insert(data, "*", result)
	return model.Result{
		Equal: result,
	}
}

func (s *Service) division(data model.Operands) model.Result {
	result := data.First / data.Second
	s.calcRepo.Insert(data, "/", result)
	return model.Result{
		Equal: result,
	}
}
