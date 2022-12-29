package repository

import "github.com/LastBit97/ewallet-restapi/model"

type TransactionRepository interface {
	CreateTransaction(*model.CreateTransactionRequest) (*model.Transaction, error)
}
