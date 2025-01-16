package calculator

import (
	"calculator/internal/model"
	"log"
)

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

func (s *Service) Subtraction(data model.Operands) model.Result {
	result := data.First - data.Second
	s.calcRepo.Insert(data, "-", result)
	return model.Result{
		Equal: result,
	}
}

func (s *Service) Multiplication(data model.Operands) model.Result {
	result := data.First * data.Second
	s.calcRepo.Insert(data, "*", result)
	return model.Result{
		Equal: result,
	}
}

func (s *Service) Division(data model.Operands) model.Result {
	result := data.First / data.Second
	if data.Second == 0 {
		log.Println("Error!")
	} else {
		s.calcRepo.Insert(data, "/", result)
		return model.Result{
			Equal: result,
		}
	}
	return model.Result{}
}

func (s *Service) Results(limit int) model.Results {
	result := s.calcRepo.Select(limit)
	res := model.Results{}
	res.List = result
	return res
}
