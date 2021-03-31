package main

import (
	"context"
	"fmt"
	"github.com/WalletService/controller"
	"github.com/WalletService/scheduler"
	"github.com/WalletService/db"
	router "github.com/WalletService/http"
	"github.com/WalletService/repository"
	"github.com/WalletService/service"
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/go-redis/redis/v8"
)

var (
	httpRouter router.IRouter
	gormDb     db.IDatabaseEngine
	gDb        *gorm.DB
	reportCron scheduler.IReportCron
)

// User
var (
	userRepository 	repository.IUserRepository
	userService		service.IUserService
	userController	controller.IUserController
)

// Wallet
var (
	walletRepository 	repository.IWalletRepository
	walletService		service.IWalletService
	walletController	controller.IWalletController
)

// Transaction
var (
	transactionRepository 	repository.ITransactionRepository
	transactionService		service.ITransactionService
	transactionController	controller.ITransactionController
)

func main() {
	httpRouter = router.NewMuxRouter()
	gormDb = db.NewGormDatabase()
	gDb = gormDb.GetDatabase()
	gormDb.RunMigration()
	initUserServiceContainer()
	initWalletServiceContainer()
	initTransactionServiceContainer()
	initCron()
	httpRouter.SERVE("8080")
}

func initUserServiceContainer() {
	userRepository = repository.NewUserRepository(gDb)
	userService = service.NewUserService(userRepository)
	userController = controller.NewUserController(userService)

	httpRouter.GET("/user/{id}", userController.GetUser)
	httpRouter.GET("/user", userController.GetUsers)
	httpRouter.POST("/user", userController.PostUser)
	httpRouter.PUT("/user/{id}", userController.PutUser)
	httpRouter.DELETE("/user/{id}", userController.DeleteUser)
}

func initWalletServiceContainer() {
	walletRepository = repository.NewWalletRepository(gDb)
	walletService = service.NewWalletService(walletRepository, userService)
	walletController = controller.NewWalletController(walletService)

	httpRouter.GET("/user/{id}/wallet", func(w http.ResponseWriter, r *http.Request) {
		walletController.GetWallet(w,r,true)
	})
	httpRouter.GET("/wallet/{id}", func(w http.ResponseWriter, r *http.Request) {
		walletController.GetWallet(w,r,false)
	})
	httpRouter.POST("/user/{id}/wallet", walletController.PostWallet)
	httpRouter.POST("/wallet/{id}/block", walletController.BlockWallet)
	httpRouter.POST("/wallet/{id}/unblock", walletController.UnBlockWallet)
}

func initTransactionServiceContainer() {
	transactionRepository = repository.NewTransactionRepository(gDb)
	transactionService = service.NewTransactionService(transactionRepository, walletService)
	transactionController = controller.NewTransactionController(transactionService)

	httpRouter.GET("/transaction", transactionController.GetTransactions)
	httpRouter.GET("/transaction/active", transactionController.GetActiveTransactions)
	httpRouter.GET("/transaction/{id}", transactionController.GetTransaction)
	httpRouter.GET("/wallet/{id}/transaction", transactionController.GetTransactionsByWalletId)
	httpRouter.POST("/wallet/{id}/transaction", transactionController.PostTransaction)
	httpRouter.PUT("/transaction/active", transactionController.UpdateActiveTransactions)
}

func initCron() {
	reportCron = scheduler.NewReportCron(transactionService)
	reportCron.StartReportCron()
	ExampleClient()
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
