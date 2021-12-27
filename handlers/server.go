package handlers

import (
	"net/http"

	"PlayerElo/elocaluclaton"

	"github.com/gorilla/mux"
)

type Server interface {
	Router() *mux.Router
}

func New(sp SiteProperties) Server {
	s := defaultServer{
		router:    mux.NewRouter(),
		siteProps: sp,
		op:        elocaluclaton.CreateRealtionalDBOperator(),
	}
	s.routes()
	return s
}

type httpMiddlewareHandler func(http.Handler) http.Handler

type SiteProperties struct {
	Title    string
	Subtitle string
}

type defaultServer struct {
	router    *mux.Router
	siteProps SiteProperties
	op        elocaluclaton.Operator
}

// Router returns the underlying router interface for the server
func (s defaultServer) Router() *mux.Router {
	return s.router
}
