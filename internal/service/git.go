package service

import (
	"github.com/asim/git-http-backend/server"
	"net/http"
	"strings"
)

func GitHttpBackend(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/git")
	server.Handler().ServeHTTP(w, r)
}
