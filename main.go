package main

import (
	"context"
	"log"
	"net/http"

	"github.com/LastBit97/ewallet-restapi/config"
	"github.com/LastBit97/ewallet-restapi/handler"
	"github.com/LastBit97/ewallet-restapi/repository"
	"github.com/LastBit97/ewallet-restapi/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	walletCollection      *mongo.Collection
	transactionCollection *mongo.Collection
	walletRepository      repository.WalletRepository
	transactionRepository repository.TransactionRepository
	walletService         service.WalletService
	transactionService    service.TransactionService
	ewalletHandler        handler.Handler
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()
	// Connect to MongoDB
	mongoOptions := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoOptions)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Println("MongoDB successfully connected...")

	walletCollection = mongoclient.Database("mongodb").Collection("wallets")
	transactionCollection = mongoclient.Database("mongodb").Collection("transactions")
	walletRepository = repository.NewWalletRepository(walletCollection, ctx)
	transactionRepository = repository.NewTransactionRepository(transactionCollection, ctx)
	walletService = service.NewDanceService(walletRepository)
	transactionService = service.NewTransactionService(transactionRepository, walletRepository)
	ewalletHandler = handler.NewHandler(walletService, transactionService)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	countWallet, err := walletCollection.EstimatedDocumentCount(context.Background())
	if err != nil {
		panic(err)
	}
	if countWallet == 0 {
		wallets, err := walletService.CreateWallets()
		if err != nil {
			log.Fatal("Failed to create wallets", err)
		}
		for _, wallet := range wallets {
			log.Printf("Create wallet with address: %s\n", wallet.Address)
		}
	}
	ewalletHandler.InitRoutes(router)
	log.Fatal(server.Run(":" + config.Port))
}
