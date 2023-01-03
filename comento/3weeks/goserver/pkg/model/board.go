package model

type Board struct {
	ID       uint   `gorm:"primarykey;autoIncrement" json:"id"`
	Email    string `json:"email"`
	First    string `json:"first"`
	Hobby    string `json:"hobby"`
	Last     string `json:"last"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
}
