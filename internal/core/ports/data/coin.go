package data

import (
	"datavisualisation/internal/core/domain"
	"io"
)

type CoinRepository interface {
	GetCoinsDataFromCSV(data io.Reader, firstLineIsTitle bool) ([]domain.Coin, error)
}
