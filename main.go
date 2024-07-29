package main

import (
	"go-stream/csv"
	"go-stream/server"
	"log"
	"net/http"
)

func main() {
	ch := csv.New("./cars_2010_2020.csv", ',')

	mux := http.NewServeMux()

	mux.HandleFunc("/large-file",
		func(w http.ResponseWriter, r *http.Request) {
			record, err := ch.Get()
			if err != nil {
				log.Panic(err)
			}
			buf := make([]byte, 0)
			for _, v := range record {
				buf = append(buf, []byte(v+"\n")...)
			}

			w.Write([]byte(buf))
		})

	s := server.New(
		server.WithAddr("localhost"),
		server.WithPort("8080"),
		server.WithMux(mux),
	)

	s.Start()
}
