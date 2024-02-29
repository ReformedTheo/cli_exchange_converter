package main

import (
	"fmt"
	"os"
)

var API_KEY string = ""

func selectCurrencies() {
	fmt.Print("")
}

func main() {
	fmt.Println(" *****Welcome to the CLI EXCHANGE CONVERTER***** \n\n\n Please input your https://www.exchangerate-api.com/  API KEY \n If you don't have one yet, you'll have to get one to use it")
	fmt.Print("API KEY: ")
	fmt.Scanln(&API_KEY)

	for {
		var option int
		fmt.Print("****MENU****\n\n 1 - Compare Currencies \n 0 - Quit app")

		switch option {
		case 1:
			selectCurrencies()
		case 0:
			os.Exit(0)
		}

	}
}
