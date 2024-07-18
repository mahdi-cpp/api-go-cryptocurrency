package models

type Order struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   string
	Coin        string
	Exchange    string
	Wallet      string
	Amount      float64
	Price       float64
	Quantity    float64 `gorm:"default:0"`
	Bitcoin     float64 `gorm:"default:43000"`
	Status      string
	Supply      uint `gorm:"default:0"`
	Market      uint `gorm:"default:0"`
	Description string
}
