package repository

import "github.com/LastBit97/ewallet-restapi/model"

type WalletRepository interface {
	CreateWallets([]*model.Wallet) ([]*model.Wallet, error)
	FindWalletByAddress(address string) (*model.Wallet, error)
	UpdateWallet(address string, data *model.Wallet) (*model.Wallet, error)
}
