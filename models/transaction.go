package models

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Ammont       float64
	Status       string
	ErrorMessage string
}

func (t *Transaction) isValid() error {
	if t.Ammont > 1000 {
		return errors.New("O valor da transação é maior que 1000")
	}
	if t.Ammont < 1 {
		return errors.New("O valor da transação deve ser maior que 1")
	}
	return nil
}
