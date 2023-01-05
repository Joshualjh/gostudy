package api

import "main.go/pkg/db"

type APIs struct {
	db *db.DBHandler
}

func NewAPI(h *db.DBHandler) *APIs {

	return &APIs{db: h}
}
