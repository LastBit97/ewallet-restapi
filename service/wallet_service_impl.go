package service

import (
	"github.com/LastBit97/ewallet-restapi/model"
	"github.com/LastBit97/ewallet-restapi/repository"
	"github.com/LastBit97/ewallet-restapi/utils"
)

type WalletServiceImpl struct {
	walletRepository repository.WalletRepository
}

func NewDanceService(walletRepository repository.WalletRepository) WalletService {
	return &WalletServiceImpl{walletRepository}
}

func (ws *WalletServiceImpl) CreateWallets() ([]*model.Wallet, error) {
	wallets := make([]*model.Wallet, 0, 10)
	for i := 0; i < cap(wallets); i++ {
		wallet := &model.Wallet{
			Address: utils.RandomString(64),
			Balance: 100,
		}
		wallets = append(wallets, wallet)
	}
	return ws.walletRepository.CreateWallets(wallets)
}
