package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type PAIR struct {
	LastUpadate    string  `json:"time_last_update_utc"`
	ConversionRate float64 `json:"conversion_rate"`
}

func pairCurrency(args []string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file", err)
	}
	API_KEY, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("API KEY NOT FOUND")
	}
	var pair PAIR
	baseCurrency := args[0]
	targetCurrency := args[1]
	url := "https://v6.exchangerate-api.com/v6/" + API_KEY + "/pair/" + baseCurrency + "/" + targetCurrency
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Could not access URL, check your API KEY", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(url)
		fmt.Println("Request failed with status code:", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	err = json.Unmarshal(body, &pair)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("The conversion rate from %s to %s is %f to 1, last update at: %s.\n", baseCurrency, targetCurrency, pair.ConversionRate, pair.LastUpadate)

}

func error(arg string) {
	log.Fatal("WRONG ARGUMENTS, CHECK HELP " + arg + ".")
}

func helpMe() {
	fmt.Print("HELP ME")
}

func registerApiKey() {
	var API_KEY string
	fmt.Print("API KEY: ")
	fmt.Scanln(&API_KEY)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	file, err := os.OpenFile(".env", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Could not open .env file ", err)
	}
	defer file.Close()

	// Check if API_KEY already exists
	scanner := bufio.NewScanner(file)
	var hasApiKey bool
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "API_KEY=") {
			hasApiKey = true
			break
		}
	}

	// If API_KEY doesn't exist, append it to the file
	if !hasApiKey {
		_, err = file.WriteString(fmt.Sprintf("API_KEY=%s\n", API_KEY))
		if err != nil {
			log.Fatal("Could not open .env file ", err)
		}
	}
}
func main() {
	args := os.Args[1:]
	switch args[0] {
	case "help":
		helpMe()
	case "register":
		registerApiKey()
	case "pair":
		if len(args) < 3 {
			error(args[0])
		}
		pairCurrency(args[1:])
	}
}
