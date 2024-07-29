package main

import (
	"go-stream/server"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/large-file",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		})

	s := server.New(
		server.WithAddr("localhost"),
		server.WithPort("8080"),
		server.WithMux(mux),
	)

	s.Start()
}
