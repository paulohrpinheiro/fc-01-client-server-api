package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	entity "github.com/paulohrpinheiro/fc-01-client-server-api/entity"
)

func GetExchange() (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", fmt.Errorf("error on http.NewRequestWithContext(): %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error on http.NewRequestWithContext(): %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error on io.ReadAll: %v", err)
	}

	var bidOutput entity.BidOutput
	err = json.Unmarshal(body, &bidOutput)
	if err != nil {
		return "", fmt.Errorf("error on json.Umarshal: %v", err)
	}

	f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	_, err = fmt.Fprintf(f, "Dólar:%v\n", bidOutput.Bid)
	if err != nil {
		log.Println(err)
	}

	return string(body), nil
}
