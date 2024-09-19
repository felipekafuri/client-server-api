package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// MUDEI PARA 500 POIS 300 NÃO FUNCIONOU
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var response map[string]string
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	bid, ok := response["bid"]
	if !ok {
		log.Fatalf("Bid not found in response")
	}

	content := fmt.Sprintf("Dólar: %s", bid)
	err = os.WriteFile("cotacao.txt", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Cotação salva em cotacao.txt")
}
