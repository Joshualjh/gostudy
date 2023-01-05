package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/pkg/model"
)

type DBHandler struct {
	gDB *gorm.DB
}

func NewConnectGorm(dsn string) (*DBHandler, error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	gormDB.AutoMigrate(&model.Students{})
	dbHandler := &DBHandler{
		gDB: gormDB,
	}
	return dbHandler, err
}
