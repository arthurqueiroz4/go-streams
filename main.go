package main

import (
	"encoding/json"
	"go-stream/csv"
	"go-stream/model"
	"go-stream/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	reponseCh := make(chan *model.Car)
	eofCh := make(chan error)

	ch := setupCsvHandler(reponseCh, eofCh)

	mux.HandleFunc("/large-file",
		func(w http.ResponseWriter, r *http.Request) {
			go ch.Chunk()

			flusher, ok := w.(http.Flusher)
			if !ok {
				http.Error(w, "Streaming not supported", http.StatusInternalServerError)
				return
			}

			for {
				select {
				case car := <-reponseCh:
					json, _ := json.Marshal(car)
					json = append(json, '\n')
					w.Write(json)
					flusher.Flush()
				case <-eofCh:
					return
				}
			}
		})

	s := server.New(
		server.WithAddr("localhost"),
		server.WithPort("8080"),
		server.WithMux(mux),
	)

	s.Start()
}

func setupCsvHandler(reponseCh chan *model.Car, eofCh chan error) *csv.CsvHandler[model.Car] {
	return &csv.CsvHandler[model.Car]{
		Filename:   "./cars_2010_2020.csv",
		Mapper:     &model.CarMapper{},
		ResponseCh: reponseCh,
		EOFCh:      eofCh,
	}
}
