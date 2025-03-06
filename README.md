# Crypto Price Tracker (Go)

This is a simple command-line cryptocurrency price tracker written in Go. It fetches real-time price data from the Binance API and displays the current price, price change, and percentage change since the program started.

## Features

* **Real-time Price Updates:** Fetches cryptocurrency prices from Binance every 5 seconds.
* **Multiple Cryptocurrency Support:** Tracks Bitcoin (BTC), Ethereum (ETH), XRP, Solana (SOL), and Dogecoin (DOGE).
* **User Selection:** Allows users to track a single cryptocurrency or all supported cryptocurrencies.
* **Color-Coded Output:** Displays price changes in green (positive) or red (negative) for easy visualization.
* **Clear Output:** Presents data in a clean and readable format.

## Prerequisites

* Internet connection.

## How to Use

**Option 1: Download Pre-built Executable (Recommended)**

1.  **Go to Releases:** Navigate to the "Releases" section of this GitHub repository.
2.  **Download Executable:** Download the appropriate executable file for your operating system (e.g., `crypto_tracker.exe` for Windows).
3.  **Run the Application:** Open a terminal or command prompt and run the downloaded executable.
4.  **Follow the Prompts:** The application will prompt you to select a cryptocurrency to track. Enter the number corresponding to your choice (1-6).

**Option 2: Build from Source (If Go is installed)**

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/ENIGMA2O5/Crypto-Prices
    cd [repository directory]
    ```
2.  **Build the Application:**
    ```bash
    go build crypto_tracker.go
    ```
3.  **Run the Application:**
    ```bash
    ./crypto_tracker
    ```
    (or `crypto_tracker.exe` on Windows)
4.  **Follow the Prompts:** The application will prompt you to select a cryptocurrency to track. Enter the number corresponding to your choice (1-6).
