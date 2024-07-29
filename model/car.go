package model

import "strconv"

type Car struct {
	Make       string  `json:"make"`
	Model      string  `json:"model"`
	FuelType   string  `json:"fuelType"`
	Year       int     `json:"year"`
	EngineSize float64 `json:"engineSize"`
	Price      float64 `json:"price"`
}

type CarMapper struct{}

func (cm *CarMapper) Map(record []string) (*Car, error) {
	year, _ := strconv.Atoi(record[3])
	engineSize, _ := strconv.ParseFloat(record[5], 64)
	price, _ := strconv.ParseFloat(record[5], 64)
	return &Car{
		Make:       record[0],
		Model:      record[1],
		FuelType:   record[3],
		Year:       year,
		EngineSize: engineSize,
		Price:      price,
	}, nil
}
