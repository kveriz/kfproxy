package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	server  *http.ServeMux
	address string
}

func New(a string) *Server {
	return &Server{
		server:  http.NewServeMux(),
		address: a,
	}
}

func (s *Server) Serve(f func(w http.ResponseWriter, r *http.Request)) {
	s.server.HandleFunc("/publish", f)
	fmt.Println("starting server...")
	if err := http.ListenAndServe(s.address, s.server); err != nil {
		log.Fatal(err)
	}
}
