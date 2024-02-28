package data

import (
	"datavisualisation/internal/core/domain"
	"datavisualisation/internal/data/mapper"
	"encoding/csv"
	"io"
)

type coinRepository struct {
	mapper *mapper.CoinMapper
}

func NewCoinRepository(mapper *mapper.CoinMapper) *coinRepository {
	return &coinRepository{
		mapper: mapper,
	}
}

func (r *coinRepository) GetCoinsDataFromCSV(data io.Reader, firstLineIsTitle bool) ([]domain.Coin, error) {
	reader := csv.NewReader(data)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var (
		coins []domain.Coin
	)

	if firstLineIsTitle {
		lines = lines[1:]
	}

	for _, line := range lines {
		coin, err := r.mapper.NewCoinFromLine(line)
		if err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}

	return coins, nil
}
