package db

import (
	"github.com/jinzhu/gorm"
	. "github.com/WalletService/model"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

type gormDatabase struct {
}

var (
	gD 			*gormDatabase
	gDOnce 		sync.Once
	gClient 	*gorm.DB
	gOnce 		sync.Once
)

// Making muxRouter instance as singleton
func NewGormDatabase() IDatabaseEngine {
	if gD == nil {
		gDOnce.Do(func() {
			gD = &gormDatabase{}
		})
	}
	return gD
}

func InitDatabase() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/WAAS?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Database connection failed : ", err)
	}else {
		log.Println("Database connection established!")
	}
	gClient = db
}

// Making sure gormClient only initialise once as singleton
func (*gormDatabase) GetDatabase() *gorm.DB {
	if gClient == nil {
		gOnce.Do(func() {
			InitDatabase()
		})
	}
	return gClient
}

func (g *gormDatabase) RunMigration() {
	if gClient == nil {
		panic("Initialise gorm db before running migrations")
	}
	gClient.AutoMigrate(&User{}, &Wallet{}, &Transaction{})

	//We need to add foreign keys manually.
	gClient.Model(&Wallet{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	gClient.Model(&Transaction{}).AddForeignKey("wallet_id", "wallets(id)", "CASCADE", "CASCADE")
}
