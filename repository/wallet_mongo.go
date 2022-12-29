package repository

import (
	"context"
	"errors"

	"github.com/LastBit97/ewallet-restapi/model"
	"github.com/LastBit97/ewallet-restapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WalletMongo struct {
	walletCollection *mongo.Collection
	ctx              context.Context
}

func NewWalletRepository(walletCollection *mongo.Collection, ctx context.Context) WalletRepository {
	return &WalletMongo{walletCollection, ctx}
}

func (wm *WalletMongo) CreateWallets(wallets []*model.Wallet) ([]*model.Wallet, error) {
	docs := make([]interface{}, len(wallets))
	for i := 0; i < len(wallets); i++ {
		docs[i] = wallets[i]
	}
	res, err := wm.walletCollection.InsertMany(wm.ctx, docs)
	if err != nil {
		return nil, err
	}

	var newWallets []*model.Wallet
	query := bson.M{"_id": bson.M{"$in": res.InsertedIDs}}
	cursor, err := wm.walletCollection.Find(wm.ctx, query)
	if err != nil {
		return nil, err
	}

	for cursor.Next(wm.ctx) {
		var wal *model.Wallet
		err := cursor.Decode(&wal)
		if err != nil {
			return nil, err
		}
		newWallets = append(newWallets, wal)
	}
	return newWallets, nil
}

func (wm *WalletMongo) FindWalletByAddress(address string) (*model.Wallet, error) {
	var wallet *model.Wallet
	query := bson.M{"address": address}
	err := wm.walletCollection.FindOne(wm.ctx, query).Decode(&wallet)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.Wallet{}, errors.New("no wallet with this address")
		}
		return nil, err
	}
	return wallet, nil
}

func (wm *WalletMongo) UpdateWallet(address string, data *model.Wallet) (*model.Wallet, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "address", Value: address}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := wm.walletCollection.FindOneAndUpdate(wm.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedWallet *model.Wallet
	if err := res.Decode(&updatedWallet); err != nil {
		return nil, errors.New("no wallet with that address exists")
	}

	return updatedWallet, nil
}
