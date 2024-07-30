package csv

import (
	"encoding/csv"
	"fmt"
	"go-stream/model"
	"os"
)

type CsvHandler[T any] struct {
	Mapper     model.DataMapper[T]
	ResponseCh chan<- *T
	EOFCh      chan<- error
	Filename   string
}

func (c *CsvHandler[T]) Chunking() *T {
	file, err := os.Open(c.Filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	r := csv.NewReader(file)
	counter := 0
	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				c.EOFCh <- err
			}
			fmt.Printf("%v files readed\n", counter)
			return nil
		}
		t, _ := c.Mapper.Map(record)

		c.ResponseCh <- t
		counter++
	}
}
