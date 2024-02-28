package main

import (
	"datavisualisation/internal/core/service"
	"datavisualisation/internal/data/mapper"
	data "datavisualisation/internal/data/repository"
	"datavisualisation/internal/http/api"
	"datavisualisation/internal/http/handler"
	"os"
)

func main() {
	filename := "coins.csv"
	file, _ := os.Open(filename)
	defer file.Close()

	r := data.NewCoinRepository(&mapper.CoinMapper{})
	s, err := service.NewCoinService(r)
	if err != nil {
		panic(err)
	}

	handler := handler.NewCoinHandler(s)
	api := api.New(handler)

	if err := api.Run(); err != nil {
		panic(err)
	}
}
