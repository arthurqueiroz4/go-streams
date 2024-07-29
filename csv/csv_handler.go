package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type CsvHandler struct {
	csvReader *csv.Reader
	file      *os.File
}

func New(filepath string, comman rune) *CsvHandler {
	f, err := os.Open(filepath)
	if err != nil {
		log.Panic(err)
	}
	c := &CsvHandler{
		csv.NewReader(f),
		f,
	}
	c.csvReader.LazyQuotes = false
	c.csvReader.Comma = comman
	return c
}

func (c *CsvHandler) Get() ([]string, error) {
	return c.csvReader.Read()
}

// TODO Insted of 'any' I can use a interface
// func (c *CsvHandler) Chunk[T any]() *T {
// 	c.csvReader.
// }
