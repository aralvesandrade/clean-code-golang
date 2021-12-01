package models

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) isValid() error {
	if t.Amount > 1000 {
		return errors.New("O valor da transação é maior que 1000")
	}
	if t.Amount < 1 {
		return errors.New("O valor da transação deve ser maior que 1")
	}
	return nil
}
