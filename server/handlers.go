package server

import (
	"io"
	"log"
	"net/http"
)

func (s *Server) homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("/ route handling")
		io.WriteString(w, "Hello")
	}
}
