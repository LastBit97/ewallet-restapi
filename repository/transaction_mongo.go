package repository

import (
	"context"

	"github.com/LastBit97/ewallet-restapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionMongo struct {
	transactionCollection *mongo.Collection
	ctx                   context.Context
}

func NewTransactionRepository(transactionCollection *mongo.Collection, ctx context.Context) TransactionRepository {
	return &TransactionMongo{transactionCollection, ctx}
}

func (tm *TransactionMongo) CreateTransaction(trans *model.CreateTransactionRequest) (*model.Transaction, error) {
	res, err := tm.transactionCollection.InsertOne(tm.ctx, trans)

	if err != nil {
		return nil, err
	}

	var newTransaction *model.Transaction
	query := bson.M{"_id": res.InsertedID}
	if err = tm.transactionCollection.FindOne(tm.ctx, query).Decode(&newTransaction); err != nil {
		return nil, err
	}

	return newTransaction, nil
}
