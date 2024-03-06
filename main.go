package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", osn)
	http.ListenAndServe(":8080", nil)
}

func osn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var URL string
		var cryptovalue string
		cprypto := r.FormValue("Cryptoval")

		if cprypto == "BTC" || cprypto == "btc" || cprypto == "Bitcoin" || cprypto == "bitcoin" {
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
			cryptovalue = "bitcoin"
		} else if cprypto == "ETH" || cprypto == "eth" || cprypto == "Ethereum" || cprypto == "ethereum" {
			URL = "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
			cryptovalue = "ethereum"
		}

		response, err := http.Get(URL)
		if err != nil {
			errMsg := fmt.Errorf("Enter information in api!: %v", err)
			fmt.Println(errMsg)
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Помилка при читанні тіла відповіді:", err)
			return
		}

		var data map[string]map[string]float64
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Помилка при розпарсуванні JSON:", err)
			return
		}
		btcUsd := data[cryptovalue]["usd"]

		Data := struct {
			Result float64
		}{
			Result: btcUsd,
		}

		tpl.Execute(w, Data)
		return
	}
	tpl.Execute(w, nil)
}
