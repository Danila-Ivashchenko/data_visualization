package service

import (
	"datavisualisation/internal/core/domain"
	"datavisualisation/internal/core/ports/data"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/go-echarts/go-echarts/v2/types"
)

type coinService struct {
	repo     data.CoinRepository
	hotCashe map[string][]domain.Coin
	dataDir  string
}

func NewCoinService(repo data.CoinRepository) (*coinService, error) {
	service := &coinService{
		repo:     repo,
		hotCashe: map[string][]domain.Coin{},
		dataDir:  "data",
	}

	err := os.Mkdir(service.dataDir, os.FileMode(0755))
	if err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	return service, nil
}

func (*coinService) generateSessionKey() string {
	id := uuid.New()
	return id.String()
}

func (s *coinService) SaveCsvFile(filename string, data io.Reader) (string, error) {
	sessionKey := s.generateSessionKey()
	os.Mkdir(fmt.Sprintf("%s/%s", s.dataDir, sessionKey), 0755)

	file, err := os.Create(fmt.Sprintf("%s/%s/%s", s.dataDir, sessionKey, filename))
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(data)

	if err != nil {
		return "", err
	}

	file.Write(bytes)
	return sessionKey, nil
}

func (s *coinService) ReadCsvData(sessionKey string, filename string, firstLineIsTitle bool) error {
	data, err := os.Open(fmt.Sprintf("%s/%s/%s", s.dataDir, sessionKey, filename))
	if err != nil {
		return err
	}

	coins, err := s.repo.GetCoinsDataFromCSV(data, firstLineIsTitle)
	if err != nil {
		return err
	}
	s.UpdateCashe(sessionKey, coins)
	return nil
}

func (s *coinService) UpdateCashe(sessionKey string, data []domain.Coin) {
	s.hotCashe[sessionKey] = data
}

func (s *coinService) GetData(sessionKey string) ([]domain.Coin, error) {
	if data, ok := s.hotCashe[sessionKey]; ok {
		return data, nil
	} else {
		files, err := os.ReadDir(fmt.Sprintf("%s/%s", s.dataDir, sessionKey))
		if err != nil && len(files) == 0 {
			return nil, fmt.Errorf("no such session key: %v", sessionKey)
		}
		if err := s.ReadCsvData(sessionKey, files[0].Name(), true); err != nil {
			return nil, fmt.Errorf("failed to read")
		}
		return s.GetData(sessionKey)
	}
}

func (s *coinService) CreateBarChart(sessionKey string, w io.Writer) error {
	bar := charts.NewBar()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Test Bar",
				Subtitle: "dogs",
			},
		),
		charts.WithYAxisOpts(
			opts.YAxis{
				Scale: true,
				AxisLabel: &opts.AxisLabel{
					Show:      true,
					FontSize:  "10",
					Formatter: "{value} dog",
					Margin:    100,
				},
			},
		),
		charts.WithXAxisOpts(
			opts.XAxis{
				Scale: true,
				AxisLabel: &opts.AxisLabel{
					Show:         true,
					Interval:     "0",
					Rotate:       30,
					ShowMinLabel: true,
					ShowMaxLabel: true,
					LineHeight:   "0",
					Margin:       10,
				},
				SplitLine: &opts.SplitLine{
					LineStyle: &opts.LineStyle{
						Type:  "solid",
						Width: 100,
						Color: "black",
					},
				},
			},
		),
	)

	data, err := s.GetData(sessionKey)
	if err != nil {
		return err
	}

	l := len(data)

	titles := make([]string, l)
	curPrice := make([]int, l)
	maxPrice := make([]int, l)

	for i, c := range data {
		titles[i] = c.Title
		curPrice[i] = c.CurrentPrice
		maxPrice[i] = c.MaxPrice
	}

	fmt.Println(titles)
	bar.SetXAxis(titles).
		AddSeries("Max price", convertData(maxPrice)).
		AddSeries("Min price", convertData(curPrice))
		// SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	return bar.Render(w)
}

func (s *coinService) CreateLineChart(sessionKey string, w io.Writer) error {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title: "Line chart",
			},
		),
		charts.WithXAxisOpts(
			opts.XAxis{
				Scale: true,
				AxisLabel: &opts.AxisLabel{
					Show:         true,
					Interval:     "0",
					Rotate:       30,
					ShowMinLabel: true,
					ShowMaxLabel: true,
					LineHeight:   "0",
				},
				SplitLine: &opts.SplitLine{
					LineStyle: &opts.LineStyle{
						Type:  "solid",
						Width: 100,
						Color: "black",
					},
				},
			},
		),
	)

	data, err := s.GetData(sessionKey)
	if err != nil {
		return err
	}

	l := len(data)

	titles := make([]string, l)
	curPrice := make([]int, l)
	maxPrice := make([]int, l)

	for i, c := range data {
		titles[i] = c.Title
		curPrice[i] = c.CurrentPrice
		maxPrice[i] = c.MaxPrice
	}

	line.SetXAxis(titles).
		AddSeries("Cur price", convertToLineData(curPrice)).
		AddSeries("Max price", convertToLineData(maxPrice))

	return line.Render(w)
}

func (s *coinService) CreateWordCloud(sessionKey string, w io.Writer) error {
	data, err := s.GetData(sessionKey)
	if err != nil {
		return err
	}

	words := map[string]int{}

	for _, c := range data {
		words[c.Title] = c.MarketCapitalization
	}

	wordCloudData := convertToWordCloudData(words)

	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title: "Coins",
			},
		),
	)

	wc.AddSeries("coins", wordCloudData)

	return wc.Render(w)
}

func (s *coinService) CreateScatter(sessionKey string, w io.Writer) error {
	data, err := s.GetData(sessionKey)
	if err != nil {
		return err
	}

	l := len(data)

	titles := make([]string, l)
	curPrice := make([]int, l)
	volatility := make([]float64, l)

	for i, c := range data {
		titles[i] = c.Title
		curPrice[i] = c.CurrentPrice
		volatility[i] = c.Volatility
	}

	scatter := charts.NewScatter()

	scatter.SetGlobalOptions(
		charts.WithYAxisOpts(
			opts.YAxis{
				Name: "Titles",
				AxisLabel: &opts.AxisLabel{
					Show:         true,
					Interval: "0",
					ShowMinLabel: true,
					ShowMaxLabel: true,

				},
			},
		),
		charts.WithXAxisOpts(
			opts.XAxis{
				Name: "Current price",
				AxisLabel: &opts.AxisLabel{
					Show:         true,
					Interval: "0",
					ShowMinLabel: true,
					ShowMaxLabel: true,
				},
			},
		),
	)

	scatter.
		SetXAxis(titles).
		AddSeries("Current price", convertToScatterData(curPrice, volatility))

	return scatter.Render(w)
}

func (s *coinService) CreatePie(sessionKey string, w io.Writer) error {
	data, err := s.GetData(sessionKey)
	if err != nil {
		return err
	}

	l := len(data)

	titles := make([]string, l)
	marketCapitalization := make([]int, l)
	allCapitalization := 0

	for i, c := range data {
		titles[i] = c.Title
		marketCapitalization[i] = c.MarketCapitalization
		allCapitalization += c.MarketCapitalization
	}

	pie := charts.NewPie()

	pie.
		AddSeries("Min Height", convertToPieData(titles, marketCapitalization))

	return pie.Render(w)
}

func convertToPieData[T any](titles []string, values []T) []opts.PieData {
	items := make([]opts.PieData, len(values))
	for i, v := range values {
		items[i] = opts.PieData{Name: titles[i], Value: v}
	}
	return items
}

func convertToScatterData[T any](values []T, size []float64) []opts.ScatterData {
	items := make([]opts.ScatterData, len(values))
	for i, v := range values {
		items[i] = opts.ScatterData{Value: v, SymbolSize: int(size[i] * 100)}
	}
	return items
}

func convertToWordCloudData[T any](data map[string]T) []opts.WordCloudData {
	items := make([]opts.WordCloudData, 0)
	for k, v := range data {
		items = append(items, opts.WordCloudData{Name: k, Value: v})
	}
	return items
}

func convertToLineData[T any](data []T) []opts.LineData {
	items := make([]opts.LineData, len(data))
	for i, v := range data {
		items[i] = opts.LineData{Value: v}
	}
	return items
}

func convertData[T any](data []T) []opts.BarData {
	items := make([]opts.BarData, len(data))
	for i, v := range data {
		items[i] = opts.BarData{Value: v}
	}
	return items
}
