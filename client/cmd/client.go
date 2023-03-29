package main

import (
	"fmt"

	clientService "github.com/paulohrpinheiro/fc-01-client-server-api/client/service"
)

func main() {
	data, err := clientService.GetExchange()
	if err != nil {
		panic(fmt.Errorf("erro ao consultar API: %v", err))
	}

	fmt.Println(data)
}
