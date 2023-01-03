package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Board struct {
	ID       uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Email    string `json:"email"`
	First    string `json:"first"`
	Hobby    string `json:"hobby"`
	Last     string `json:"last"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
}

type DBHandler struct {
	gDB *gorm.DB
}

func main() {

	dbHandler, err := NewAndConnectGorm("cruduser:sADMIN123@tcp(crudserverset.mysql.database.azure.com)/dev?charset=utf8mb4&parseTime=True&loc=Local&tls=true")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(dbHandler)
}

func NewAndConnectGorm(dsn string) (*DBHandler, error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	gormDB.AutoMigrate(Board{})
	dbHandler := &DBHandler{
		gDB: gormDB,
	}

	return dbHandler, err
}
