package csv

import "strconv"

// TODO Move to "model" and create a interface Mapable -> MapFrom([]string)
type Car struct {
	Make       string
	Model      string
	FuelType   string
	Year       int
	EngineSize float64
	Price      float64
}

func MapFrom(record []string) *Car {
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
	}
}
