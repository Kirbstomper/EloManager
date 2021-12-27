package handlers

import "net/http"

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func (s *defaultServer) routes() {
	s.router.PathPrefix("/js/").HandlerFunc(s.serveStaticResource()).Methods(http.MethodGet)
	s.router.PathPrefix("/").HandlerFunc(s.serveIndexPage()).Methods(http.MethodGet)
	s.router.PathPrefix("/add").HandlerFunc(s.addPlayers()).Methods(http.MethodPost)
	s.router.PathPrefix("/decide").HandlerFunc(s.decide()).Methods(http.MethodPost)
	s.router.PathPrefix("/updateElo").HandlerFunc(s.updatePlayerElo()).Methods(http.MethodPost)
	s.router.PathPrefix("/getPlayer").HandlerFunc(s.retrievePlayerInformation()).Methods(http.MethodGet)
}
