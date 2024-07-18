package repository

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-cryptocurrency/config"
	"github.com/mahdi-cpp/api-go-cryptocurrency/email"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ctx = context.Background()

type Hotbit struct {
	Error  interface{} `json:"error"`
	Result string      `json:"result"`
	Id     int         `json:"id"`
}

type PancakeSwap struct {
	NativePrice struct {
		Value    string `json:"value"`
		Decimals int    `json:"decimals"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
	} `json:"nativePrice"`
	UsdPrice        float64 `json:"usdPrice"`
	ExchangeAddress string  `json:"exchangeAddress"`
	ExchangeName    string  `json:"exchangeName"`
}

func Wojak() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var loopDelay = time.Second * 10
	var emailDelay = time.Minute * 10

	for {
		s := send(client, "https://deep-index.moralis.io/api/v2/erc20/0x55f96c7005d7c684a65ee653b07b5fe1507c56ab/price?chain=bsc", "N36kTr2Qg8HWW46IRZ3rO1oTxguAXQrqJeVK0F5TzOWWYeZrbrNbDH9fbKiZOgwG")

		var PancakeSwap PancakeSwap
		err := json.Unmarshal([]byte(s), &PancakeSwap)
		if err != nil {
			fmt.Printf("Error occured during unmarshaling. Error: %s\n", err.Error())
			fmt.Println("PancakeSwap server did not return a valid JSON response")
			time.Sleep(time.Second * 6)
			continue
		}

		s = send(client, "https://api.hotbit.io/api/v1/market.last?market=WOJ/USDT", "")
		var Hotbit Hotbit
		err = json.Unmarshal([]byte(s), &Hotbit)
		if err != nil {
			fmt.Printf("Error occured during unmarshaling. Error: %s\n", err.Error())
			fmt.Println("Hotbit server did not return a valid JSON response")
			time.Sleep(time.Second * 6)
			continue
		}

		PaSwapPrice := fmt.Sprintf("%.3f", PancakeSwap.UsdPrice)
		HotbitPrice, err := strconv.ParseFloat(Hotbit.Result, 64)
		fmt.Println("PaSwap price:", PaSwapPrice)
		fmt.Println("Hotbit Price:", fmt.Sprintf("%.3f", HotbitPrice))

		var divide = 0.0
		if PancakeSwap.UsdPrice > HotbitPrice {
			divide = PancakeSwap.UsdPrice / HotbitPrice
			_, fractional := math.Modf(divide) // get fractional of divide
			var profit = int(fractional * 100) //calculate profit percent of this arbitrage
			fmt.Println("profit:", profit, "--> pancakeswap is more than hotbit, required more than 14%")

			if profit > 14 {
				_, err := config.Redis.Get(ctx, "email").Result()
				if err == nil { //email is exit!, wait until email expiration
					time.Sleep(loopDelay)
					continue
				}

				err = config.Redis.Set(ctx, "email", "mahdi.cpp@gmail.com", emailDelay).Err()
				if err != nil {
					fmt.Println(err)
				}

				message := `<b>` + "PancakeSwap is more than Hotbit" + `</b>` +
					`<p><b>` + "Arbitrage is more of 14%" + `</b></p>` +
					`<p><b>Pancakeswap: ` + fmt.Sprintf("%f", PancakeSwap.UsdPrice) + `</b></p>` +
					`<p><b>Hotbit     : ` + fmt.Sprintf("%f", HotbitPrice) + `</b></p>`
				email.SendEmail("Wojak Arbitrage", message)
			}
		} else {
			divide = HotbitPrice / PancakeSwap.UsdPrice
			_, fractional := math.Modf(divide) // get fractional of divide
			var profit = int(fractional * 100) //calculate profit percent of this arbitrage
			fmt.Println("profit:", profit, "--> hotbit is more than pancakeswap, required more than 15%")
			if profit > 15 {
				_, err := config.Redis.Get(ctx, "email").Result()
				if err == nil {
					time.Sleep(loopDelay)
					continue
				}

				err = config.Redis.Set(ctx, "email", "mahdi.cpp@gmail.com", emailDelay).Err()
				if err != nil {
					fmt.Println(err)
				}

				message := `<b>` + "Hotbit is more than PancakeSawp" + `</b>` +
					`<p><b>` + "Arbitrage is more of 15%" + `</b></p>` +
					`<p><b>Pancakeswap: ` + fmt.Sprintf("%f", PancakeSwap.UsdPrice) + `</b></p>` +
					`<p><b>Hotbit     : ` + fmt.Sprintf("%f", HotbitPrice) + `</b></p>`
				email.SendEmail("Wojak Arbitrage", message)
			}
		}

		fmt.Println("")
		time.Sleep(loopDelay)
	}
}

func send(client *http.Client, url string, header string) string {
	req, _ := http.NewRequest("GET", url, nil)

	if len(header) > 0 {
		req.Header.Set("X-API-Key", header)
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	//defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	s := strings.TrimSpace(string(content))
	return s
}
