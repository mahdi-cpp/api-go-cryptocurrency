package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-cryptocurrency/models"
	"github.com/mahdi-cpp/api-go-cryptocurrency/repository"
	"math"
	"strconv"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func str2uint64(str string) (uint, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	return uint(i), err
}

func AddCryptoRoutes(rg *gin.RouterGroup) {

	router := rg.Group("/blockchain")

	router.GET("/getAll", func(c *gin.Context) {
		//c.JSON(200, repository.GetAllTutorials())
		c.JSON(200, repository.GetAllUser2())
	})

	router.GET("/getByCrypto", func(c *gin.Context) {
		c.JSON(200, repository.GetUser2ByCrypto())
	})

	router.GET("/getAllOrders", func(c *gin.Context) {
		c.JSON(200, repository.GetAllOrders())
	})

	router.POST("/create", func(c *gin.Context) {
		order := models.Order{}
		s := c.Query("data")
		json.Unmarshal([]byte(s), &order)
		fmt.Println(order.Wallet)
		order.Quantity = toFixed(order.Quantity, 6)
		repository.CreateOrder(order)
		c.JSON(200, "Ok Mahdi")
	})

	router.POST("/setCrypto", func(c *gin.Context) {
		fmt.Println(c.Query("crypto"))
		repository.SetCrypto(c.Query("crypto"))
	})
	router.POST("/remove", func(c *gin.Context) {
		var str = c.Query("id")
		id, _ := str2uint64(str)
		fmt.Println("remove id: ", id)
		repository.RemoveOrder(id)
	})
}
