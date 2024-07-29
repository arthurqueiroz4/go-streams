package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type CsvHandler struct {
	csvReader csv.Reader
	comman    rune
}

func New(filepath string, comman rune) *CsvHandler {
	f, err := os.Open(filepath)
	if err != nil {
		log.Panic(err)
	}

	return &CsvHandler{
		*csv.NewReader(f),
		comman,
	}
}
