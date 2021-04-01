package db

import (
	"github.com/jinzhu/gorm"
	. "github.com/WalletService/model"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

type gormDatabase struct {
	client 	*gorm.DB
	once 	sync.Once
}

func NewGormDatabase() IDatabaseEngine {
	return &gormDatabase{}
}

func InitDatabase(g *gormDatabase) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/WAAS?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Database connection failed : ", err)
	}else {
		log.Println("Database connection established!")
	}
	log.Println("MySql connection running on port 3306")
	g.client = db
}

// Making sure gormClient only initialise once as singleton
func (g *gormDatabase) GetDatabase() *gorm.DB {
	if g.client == nil {
		g.once.Do(func() {
			InitDatabase(g)
		})
	}
	return g.client
}

func (g *gormDatabase) RunMigration() {
	if g.client == nil {
		panic("Initialise gorm db before running migrations")
	}
	g.client.AutoMigrate(&User{}, &Wallet{}, &Transaction{})

	//We need to add foreign keys manually.
	g.client.Model(&Wallet{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	g.client.Model(&Transaction{}).AddForeignKey("wallet_id", "wallets(id)", "CASCADE", "CASCADE")
}
