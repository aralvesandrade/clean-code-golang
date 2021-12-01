package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	transaction := Transaction{}
	transaction.ID = "1"
	transaction.AccountID = "1"

	t.Run("Teste deve executar a validação sem erro", func(t *testing.T) {
		transaction.Ammont = 200
		isValid := transaction.isValid()
		assert.Nil(t, isValid)
	})

	t.Run("Teste deve executar a validação com erro de transação maior que 1000", func(t *testing.T) {
		transaction.Ammont = 2000
		isValid := transaction.isValid()
		assert.NotNil(t, isValid)
		assert.Equal(t, isValid.Error(), "O valor da transação é maior que 1000")
	})

	t.Run("Teste deve executar a validação com erro de transação deve ser maior que 1", func(t *testing.T) {
		transaction.Ammont = 0
		isValid := transaction.isValid()
		assert.NotNil(t, isValid)
		assert.Equal(t, isValid.Error(), "O valor da transação deve ser maior que 1")
	})
}
