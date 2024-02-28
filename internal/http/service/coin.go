package service

import (
	"datavisualisation/internal/core/domain"
	"io"
)

type CoinService interface {
	CreateBarChart(sessionKey string, w io.Writer) error
	CreateLineChart(sessionKey string, w io.Writer) error
	CreatePie(sessionKey string, w io.Writer) error
	CreateScatter(sessionKey string, w io.Writer) error
	CreateWordCloud(sessionKey string, w io.Writer) error

	ReadCsvData(sessionKey string, filename string, firstLineIsTitle bool) error
	SaveCsvFile(filename string, data io.Reader) (string, error)
	UpdateCashe(sessionKey string, data []domain.Coin)
}
