package mapper

import (
	"datavisualisation/internal/core/domain"
	"strconv"
	"time"
)

type CoinMapper struct{}

func (*CoinMapper) NewCoinFromLine(line []string) (domain.Coin, error) {
	coin := domain.Coin{}

	date, err := time.Parse("2006-01-02", line[1])
	if err != nil {
		return coin, err
	}
	cPrice, err := strconv.Atoi(line[2])
	if err != nil {
		return coin, err
	}
	mPrice, err := strconv.Atoi(line[3])
	if err != nil {
		return coin, err
	}
	capitalization, err := strconv.Atoi(line[4])
	if err != nil {
		return coin, err
	}
	volatility, err := strconv.ParseFloat(line[5], 64)
	if err != nil {
		return coin, err
	}
	return domain.Coin{
		Title:                line[0],
		AppearanceDate:       date,
		CurrentPrice:         cPrice,
		MaxPrice:             mPrice,
		MarketCapitalization: capitalization,
		Volatility:           volatility,
	}, nil
}
