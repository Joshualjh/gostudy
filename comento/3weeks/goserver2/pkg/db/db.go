package db

import "gorm.io/gorm"

type DBHandler struct {
	gDB *gorm.DB
}
