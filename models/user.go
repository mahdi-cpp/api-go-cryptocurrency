package models

type User2 struct {
	ID         uint   `json:"id",gorm:"primarykey"`
	Name       string `json:"name"`
	AvatarUrl  string `json:"avatarUrl"`
	Title      string `json:"title"`
	Crypto     string `json:"crypto"`
	Company    string `json:"company"`
	IsVerified bool   `json:"isVerified"`
	Status     string `json:"status"`
	Role       string `json:"role"`
}
