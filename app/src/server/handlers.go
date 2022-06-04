package server

import "net/http"

func (s *Server) homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	}
}
