package service

import (
	"errors"
	"time"

	"github.com/LastBit97/ewallet-restapi/model"
	"github.com/LastBit97/ewallet-restapi/repository"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	walletRepository      repository.WalletRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, walletRepository repository.WalletRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepository, walletRepository}
}

func (ts *TransactionServiceImpl) Send(trans *model.CreateTransactionRequest) (*model.Transaction, error) {
	walletFrom, err := ts.walletRepository.FindWalletByAddress(trans.AddressFrom)
	if err != nil {
		return nil, err
	}
	if walletFrom.Balance-trans.Amount < 0 {
		return nil, errors.New("insufficient funds")
	}
	walletTo, err := ts.walletRepository.FindWalletByAddress(trans.AddressTo)
	if err != nil {
		return nil, err
	}
	walletFrom.Balance -= trans.Amount
	walletTo.Balance += trans.Amount
	_, err = ts.walletRepository.UpdateWallet(walletFrom.Address, walletFrom)
	if err != nil {
		return nil, err
	}
	_, err = ts.walletRepository.UpdateWallet(walletTo.Address, walletTo)
	if err != nil {
		return nil, err
	}

	trans.CreateAt = time.Now()
	newTransaction, err := ts.transactionRepository.CreateTransaction(trans)
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}
