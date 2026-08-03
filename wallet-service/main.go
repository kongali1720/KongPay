package main

import (
	"fmt"
	"net/http"
)

func getBalance(w http.ResponseWriter, r *http.Request) {
	// Nanti di sini kita akan ambil data dari PostgreSQL
	fmt.Fprintf(w, "{\"user_id\": \"user123\", \"balance\": 1000000000.00}")
}

func main() {
	http.HandleFunc("/api/v1/wallet/balance", getBalance)
	fmt.Println("Wallet Service berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)
}
