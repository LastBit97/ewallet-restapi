package service

import "github.com/LastBit97/ewallet-restapi/model"

type WalletService interface {
	CreateWallets() ([]*model.Wallet, error)
	GetWallet(address string) (*model.Wallet, error)
}
