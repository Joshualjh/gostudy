package model

type Students struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Rollno string `gorm:"primarykey" json:"rollno"`
}
