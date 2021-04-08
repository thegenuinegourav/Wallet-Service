package main

import (
	"encoding/json"
	"github.com/WalletService/cache"
	config "github.com/WalletService/config"
	"github.com/WalletService/controller"
	"github.com/WalletService/db"
	_ "github.com/WalletService/docs" // This line is necessary for go-swagger to find your docs!
	router "github.com/WalletService/http"
	"github.com/WalletService/repository"
	"github.com/WalletService/scheduler"
	"github.com/WalletService/service"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
)

var (
	c 		   config.Config
	httpRouter router.IRouter
	gormDb     db.IDatabaseEngine
	gDb        *gorm.DB
)

// Cron
var (
	cronCache  cache.ICronCache
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
	walletCache         cache.IWalletCache
	walletRepository 	repository.IWalletRepository
	walletService		service.IWalletService
	walletController	controller.IWalletController
)

// Transaction
var (
	transactionIdempotentCache cache.ITransactionIdempotentCache
	transactionRepository 	repository.ITransactionRepository
	transactionService		service.ITransactionService
	transactionController	controller.ITransactionController
)

func main() {
	initConfig()
	httpRouter = router.NewMuxRouter()
	httpRouter.ADDVERSION("/api/v1")
	gormDb = db.NewGormDatabase()
	gDb = gormDb.GetDatabase(c.Database)
	gormDb.RunMigration()
	initCachingLayer()
	initUserServiceContainer()
	initWalletServiceContainer()
	initTransactionServiceContainer()
	initCron()
	httpRouter.SERVE(c.App.Port)
}

func initConfig() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Printf("No ./config.json file found!! Terminating the server, error: %s\n", err.Error())
		panic("No config file found! Error : " + err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Printf("Error occurred while decoding json to config model, error: %s\n", err.Error())
		panic(err.Error())
	}
}

func initCachingLayer() {
	cacheEngine := cache.NewCache(c.Cache, c.Cache.Wallet.Db)
	cacheEngine.GetCacheClient()
	if err := cacheEngine.CheckConnection(); err != nil {
		panic(err)
	}
	walletCache = cache.NewWalletCache(cacheEngine, c.Cache.Wallet)		// 0 means no expiry date

	cacheEngine2 := cache.NewCache(c.Cache, c.Cache.Idempotent.Db)
	cacheEngine2.GetCacheClient()
	if err := cacheEngine2.CheckConnection(); err != nil {
		panic(err)
	}
	transactionIdempotentCache = cache.NewTransactionIdempotentCache(cacheEngine2, c.Cache.Idempotent)

	cacheEngine3 := cache.NewCache(c.Cache, c.Cache.CronLock.Db)
	cacheEngine3.GetCacheClient()
	if err := cacheEngine3.CheckConnection(); err != nil {
		panic(err)
	}
	cronCache = cache.NewCronCache(cacheEngine3, c.Cache.CronLock)
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
	walletService = service.NewWalletService(walletRepository, userService, walletCache)
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
	transactionController = controller.NewTransactionController(transactionService, transactionIdempotentCache)

	httpRouter.GET("/transaction", transactionController.GetTransactions)
	httpRouter.GET("/transaction/active", transactionController.GetActiveTransactions)
	httpRouter.GET("/transaction/{id}", transactionController.GetTransaction)
	httpRouter.GET("/wallet/{id}/transaction", transactionController.GetTransactionsByWalletId)
	httpRouter.POST("/wallet/{id}/transaction", transactionController.PostTransaction)
	httpRouter.PUT("/transaction/active", transactionController.UpdateActiveTransactions)
}

func initCron() {
	reportCron = scheduler.NewReportCron(transactionService, cronCache)
	reportCron.StartReportCron()
}