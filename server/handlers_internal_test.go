package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rtnhnd255/osau_stud_org_site.git/storage"
	"github.com/stretchr/testify/assert"
)

func baseServer() *Server {
	storc, _ := storage.NewConfig("./test_server_config.json")
	serc, _ := NewConfig("./test_config.json", storc)
	return NewServer(serc)
}
func TestHomeHandler(t *testing.T) {
	s := baseServer()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.homeHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
