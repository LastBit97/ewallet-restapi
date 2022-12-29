package repository

import (
	"context"

	"github.com/LastBit97/ewallet-restapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (tm *TransactionMongo) GetTransactions(limit int) ([]*model.Transaction, error) {
	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{}

	cursor, err := tm.transactionCollection.Find(tm.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(tm.ctx)

	var transactions []*model.Transaction

	for cursor.Next(tm.ctx) {
		transaction := &model.Transaction{}
		err := cursor.Decode(transaction)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return []*model.Transaction{}, nil
	}

	return transactions, nil
}
