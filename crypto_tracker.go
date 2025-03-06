package main

import (
        "encoding/json"
        "fmt"
        "io"
        "log"
        "net/http"
        "strconv"
        "strings"
        "time"
)

// BinanceTickerPrice represents the JSON structure from Binance API
type BinanceTickerPrice struct {
        Symbol string `json:"symbol"`
        Price  string `json:"price"`
}

// getCryptoPrice fetches cryptocurrency price from Binance API
func getCryptoPrice(symbol string) (float64, error) {
        url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
        resp, err := http.Get(url)
        if err != nil {
                return 0, err
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
                return 0, fmt.Errorf("HTTP error: %s", resp.Status)
        }

        body, err := io.ReadAll(resp.Body)
        if err != nil {
                return 0, err
        }

        var ticker BinanceTickerPrice
        err = json.Unmarshal(body, &ticker)
        if err != nil {
                return 0, err
        }

        price, err := strconv.ParseFloat(ticker.Price, 64)
        if err != nil {
                return 0, err
        }
        return price, nil
}

// displayCryptoData displays cryptocurrency data with color-coded price change
func displayCryptoData(cryptoName, symbol string, initialPrice float64) (float64, error) {
        currentPrice, err := getCryptoPrice(symbol)
        if err != nil {
                return 0, err
        }

        priceChange := currentPrice - initialPrice
        percentageChange := (priceChange / initialPrice) * 100

        color := "\033[32m" // Green
        if priceChange < 0 {
                color = "\033[31m" // Red
        }
        resetColor := "\033[0m"

        fmt.Printf("\n%s Price: $%.4f\n", cryptoName, currentPrice)
        fmt.Printf("Price Change Since Start: %s$%.4f (%.4f%%)%s\n", color, priceChange, percentageChange, resetColor)

        return currentPrice, nil
}

func main() {
        fmt.Println("Select a cryptocurrency to track:")
        fmt.Println("1. Bitcoin (BTC)")
        fmt.Println("2. Ethereum (ETH)")
        fmt.Println("3. XRP")
        fmt.Println("4. Solana (SOL)")
        fmt.Println("5. Dogecoin (DOGE)")
        fmt.Println("6. Track All (BTC, ETH, XRP, SOL, DOGE)")

        var choice string
        fmt.Print("Enter the number corresponding to your choice (1, 2, 3, 4, 5, or 6): ")
        fmt.Scanln(&choice)
        choice = strings.TrimSpace(choice)

        var symbols []struct {
                Name   string
                Symbol string
        }

        switch choice {
        case "1":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{{"Bitcoin", "BTCUSDT"}}
        case "2":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{{"Ethereum", "ETHUSDT"}}
        case "3":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{{"XRP", "XRPUSDT"}}
        case "4":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{{"Solana", "SOLUSDT"}}
        case "5":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{{"Dogecoin", "DOGEUSDT"}}
        case "6":
                symbols = []struct {
                        Name   string
                        Symbol string
                }{
                        {"Bitcoin", "BTCUSDT"},
                        {"Ethereum", "ETHUSDT"},
                        {"XRP", "XRPUSDT"},
                        {"Solana", "SOLUSDT"},
                        {"Dogecoin", "DOGEUSDT"},
                }
        default:
                fmt.Println("Invalid choice. Exiting.")
                return
        }

        initialPrices := make(map[string]float64)
        for _, s := range symbols {
                initialPrice, err := getCryptoPrice(s.Symbol)
                if err != nil {
                        fmt.Printf("Failed to fetch initial %s price: %v. Skipping.\n", s.Name, err)
                        continue
                }
                initialPrices[s.Symbol] = initialPrice
                fmt.Printf("\nInitial %s Price: $%.4f\n", s.Name, initialPrice)
        }

        for {
                time.Sleep(5 * time.Second)
                fmt.Println("\n" + strings.Repeat("=", 50))
                for _, s := range symbols {
                        if _, ok := initialPrices[s.Symbol]; !ok {
                                continue
                        }
                        currentPrice, err := displayCryptoData(s.Name, s.Symbol, initialPrices[s.Symbol])
                        if err != nil {
                                log.Printf("Error displaying %s data: %v", s.Name, err)
                                continue
                        }
                        initialPrices[s.Symbol] = currentPrice
                }
        }
}