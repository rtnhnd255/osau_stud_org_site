package server

import "github.com/gorilla/mux"

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (*Server) Start() error {
	return nil
}
