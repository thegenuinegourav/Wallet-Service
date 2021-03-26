package db

import "github.com/jinzhu/gorm"

type IDatabaseEngine interface {
	// TODO read from config file
	GetDatabase() *gorm.DB
	RunMigration()
}
