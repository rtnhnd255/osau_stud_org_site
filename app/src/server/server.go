package server

import "github.com/gorilla/mux"

type Server struct {
	router *mux.Router
	port   string
}

func NewServer(cfg *Config) *Server {
	res := &Server{
		router: mux.NewRouter(),
		port:   cfg.Port,
	}
	res.configureRouter()
	return res
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.homeHandler())
}

func (s *Server) Run() error {
	return nil
}
