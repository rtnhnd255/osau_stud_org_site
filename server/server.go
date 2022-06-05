package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtnhnd255/osau_stud_org_site.git/storage"
)

type Server struct {
	config *Config
	store  *storage.Storage
	router *mux.Router
	addr   string
}

func NewServer(cfg *Config) *Server {
	log.Println("Configuring server")

	res := &Server{
		router: mux.NewRouter(),
		addr:   cfg.Addr,
		config: cfg,
		store:  &storage.Storage{},
	}
	res.configureRouter()

	if err := res.configureStorage(); err != nil {
		log.Println("Trouble with configuring storage")
		log.Fatal(err)
	}
	return res
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.homeHandler())
}

func (s *Server) configureStorage() error {
	_s := storage.New(s.config.Store)

	if err := _s.Open(); err != nil {
		log.Println("Trouble with test opening storage")
		return err
	}

	s.store = _s
	return nil
}

func (s *Server) Run() error {
	log.Println("Starting server")
	return http.ListenAndServe(s.addr, s.router)
}
