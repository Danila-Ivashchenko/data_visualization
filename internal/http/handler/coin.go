package handler

import "datavisualisation/internal/http/service"

type coinHandler struct {
	service service.CoinService
}

func NewCoinHandler(service service.CoinService) *coinHandler {
	return &coinHandler{service: service}
}
