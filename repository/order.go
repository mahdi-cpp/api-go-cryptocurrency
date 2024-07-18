package repository

import (
	"github.com/mahdi-cpp/api-go-cryptocurrency/config"
	"github.com/mahdi-cpp/api-go-cryptocurrency/models"
)

func CreateOrder(order models.Order) {
	config.DB.Debug().Create(&order)
}
func RemoveOrder(id uint) {
	config.DB.Debug().Delete(&models.Order{}, id)
}

func GetAllOrders() []models.Order {
	var orders []models.Order
	config.DB.Debug().Order("created_at DESC").Find(&orders)
	return orders
}

func GetAllUser2() []models.User2 {
	var users []models.User2
	config.DB.Debug().Limit(30).Find(&users)
	return users
}

func GetUser2ByCrypto() []models.User2 {
	var users []models.User2
	config.DB.Debug().Where("crypto = ?", "Bitcoin").Find(&users)
	return users
}

func SetCrypto(crypto string) {
	config.DB.Debug().Model(models.User2{}).Where("id = ?", 1).Update("crypto", crypto)
}
