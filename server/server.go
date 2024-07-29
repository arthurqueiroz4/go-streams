package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	mux  *http.ServeMux
	addr string
	port string
}

type Option func(*Server)

func New(opts ...Option) *Server {
	s := &Server{
		addr: "localhost",
		port: ":3000",
		mux:  http.NewServeMux(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

func WithPort(port string) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithMux(mux *http.ServeMux) Option {
	return func(s *Server) {
		s.mux = mux
	}
}

func (s *Server) AddHandler(uri string, handler http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(uri, handler)
}

func (s *Server) Start() {
	http := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.addr, s.port),
		Handler: s.mux,
	}

	log.Printf("Listening port: %s", s.port)
	http.ListenAndServe()
}
