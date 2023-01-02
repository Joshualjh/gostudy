package api

import "main/pkg/db"

type APIs struct {
	db *db.DBHandler
}

func NewAPI(h *db.DBHandler) *APIs {
	return &APIs{db: h}
}
