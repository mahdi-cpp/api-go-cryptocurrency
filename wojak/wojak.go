package wojak

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Crypto struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

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

func GetCrypto() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	s := send(client, "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc", "")
	var crypto []Crypto
	err := json.Unmarshal([]byte(s), &crypto)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	} else {
		for _, element := range crypto {
			//fmt.Println(index+1, ":", element.ID)
			//fmt.Println("{ label: '", i, "', value: '", element.Name, "'},")
			fmt.Print("'", element.Name, "',")
		}
	}
}

func Wojak() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	for {
		s := send(client, "https://deep-index.moralis.io/api/v2/erc20/0x55f96c7005d7c684a65ee653b07b5fe1507c56ab/price?chain=bsc", "N36kTr2Qg8HWW46IRZ3rO1oTxguAXQrqJeVK0F5TzOWWYeZrbrNbDH9fbKiZOgwG")
		//fmt.Println(s)

		var PancakeSwap PancakeSwap
		err := json.Unmarshal([]byte(s), &PancakeSwap)
		if err != nil {
			log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
		}

		s = send(client, "https://api.hotbit.io/api/v1/market.last?market=WOJ/USDT", "")
		var Hotbit Hotbit
		err = json.Unmarshal([]byte(s), &Hotbit)
		if err != nil {
			log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
		}

		PaSwapPrice := fmt.Sprintf("%.3f", PancakeSwap.UsdPrice)
		HotbitPrice, err := strconv.ParseFloat(Hotbit.Result, 64)

		var priceA = PancakeSwap.UsdPrice
		var priceB = HotbitPrice
		var divide = 0.0

		if priceA > priceB {
			divide = priceA / priceB
		} else {
			divide = priceB / priceA
		}
		_, fractional := math.Modf(divide) // get fractional of divide
		var profit = int(fractional * 100) //calculate profit percent of this arbitrage

		fmt.Println("PaSwap price:", PaSwapPrice)
		fmt.Println("Hotbit Price:", fmt.Sprintf("%.3f", HotbitPrice))
		fmt.Println("Profit Percent:", profit)

		fmt.Println("")
		time.Sleep(time.Second * 6)
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
