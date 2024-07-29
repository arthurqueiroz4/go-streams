package csv

import (
	"encoding/csv"
	"go-stream/model"
	"os"
)

type CsvHandler[T any] struct {
	Mapper     model.DataMapper[T]
	ResponseCh chan<- *T
	EOFCh      chan<- error
	Filename   string
}

func (c *CsvHandler[T]) Chunk() *T {
	file, err := os.Open(c.Filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				c.EOFCh <- err
			}
			return nil
		}
		t, _ := c.Mapper.Map(record)

		c.ResponseCh <- t
	}
}
