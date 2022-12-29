package service

import "github.com/LastBit97/ewallet-restapi/model"

type TransactionService interface {
	Send(*model.CreateTransactionRequest) (*model.Transaction, error)
	GetTransactions(count int) ([]*model.Transaction, error)
}
