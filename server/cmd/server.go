package main

import (
	"log"
	"net/http"

	serverDb "github.com/paulohrpinheiro/fc-01-client-server-api/server/infra"
	serverService "github.com/paulohrpinheiro/fc-01-client-server-api/server/service"
)

func main() {
	serverDb.Create()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", GetExchangeHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func GetExchangeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "only GET for GetExchangeHandler endpoint", http.StatusNotImplemented)
	}

	res, err := serverService.GetExchangeFromAPI()
	if err != nil {
		http.Error(w, "error on get data", http.StatusInternalServerError)
	}

	w.Write(res)
}
