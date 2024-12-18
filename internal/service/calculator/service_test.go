package calculator

import (
	"calculator/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
)

func TestServer_Add(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCalculatorRepository := NewMockCalculatorRepository(ctrl)

	t.Run("two_positive_ok", func(t *testing.T) {
		data := model.Operands{
			First:  rand.Float64() * 1000,
			Second: rand.Float64() * 100,
		}
		expextedResult := data.First + data.Second
		log.Println(data.First, data.Second, expextedResult)

		mockCalculatorRepository.EXPECT().Insert(data, "+", expextedResult)

		srv := New(mockCalculatorRepository)
		res := srv.Add(data)

		assert.Equal(t, expextedResult, res.Equal)
	})

	t.Run("two_positive_ok_with_zero", func(t *testing.T) {
		data := model.Operands{
			First:  rand.Float64() * 1000,
			Second: 0,
		}
		expextedResult := data.First + data.Second
		log.Println(data.First, data.Second, expextedResult)

		mockCalculatorRepository.EXPECT().Insert(data, "+", expextedResult)

		srv := New(mockCalculatorRepository)
		res := srv.Add(data)

		assert.Equal(t, expextedResult, res.Equal)
	})
}
