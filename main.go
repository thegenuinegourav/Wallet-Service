package main

import (
	"github.com/WalletService/controller"
	"github.com/WalletService/cron"
	"github.com/WalletService/db"
	router "github.com/WalletService/http"
	"github.com/WalletService/repository"
	"github.com/WalletService/service"
	"github.com/jinzhu/gorm"
	"net/http"
)

var (
	httpRouter 	router.IRouter
	gormDb		db.IDatabaseEngine
	gDb			*gorm.DB
	reportCron  cron.IReportCron
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
	reportCron = cron.NewReportCron(transactionService)
	reportCron.StartReportCron()
}
