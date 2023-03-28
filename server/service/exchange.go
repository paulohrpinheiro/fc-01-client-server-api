package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	entity "github.com/paulohrpinheiro/fc-01-client-server-api/entity"
	db "github.com/paulohrpinheiro/fc-01-client-server-api/server/infra"
)

func getFromApi() ([]byte, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, fmt.Errorf("error on http.NewRequestWithContext(): %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error on http.NewRequestWithContext(): %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error on io.ReadAll: %v", err)
	}

	return body, nil
}

func logToDb(bid string) {
	err := db.Write(bid)
	if err != nil {
		log.Printf("Error saving to database: %v", err)
	}
}

func GetExchangeFromAPI() ([]byte, error) {
	var bidInput entity.BidInput

	body, err := getFromApi()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &bidInput)
	if err != nil {
		return nil, fmt.Errorf("error on json.Umarshal: %v", err)
	}

	bidOutput := entity.BidOutput{Bid: bidInput.Usdbrl.Bid}

	jsonBidOutput, err := json.Marshal(bidOutput)
	if err != nil {
		return nil, fmt.Errorf("error on json.Marshal: %v", err)
	}

	logToDb(bidInput.Usdbrl.Bid)
	return jsonBidOutput, nil
}
