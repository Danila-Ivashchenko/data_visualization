package domain

import "time"

type Coin struct {
	Title                   string
	AppearanceDate time.Time
	CurrentPrice           int
	MaxPrice               int
	MarketCapitalization   int
	Volatility             float64
}
