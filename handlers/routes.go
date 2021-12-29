package handlers

import "net/http"

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func (s *defaultServer) routes() {
	s.router.PathPrefix("/js/").HandlerFunc(s.serveStaticResource()).Methods(http.MethodGet)
	s.router.Path("/").HandlerFunc(s.serveIndexPage()).Methods(http.MethodGet)
	s.router.Path("/add").HandlerFunc(s.addPlayers()).Methods(http.MethodPost)
	s.router.Path("/decide").HandlerFunc(s.decide()).Methods(http.MethodPost)
	s.router.Path("/updateElo").HandlerFunc(s.updatePlayerElo()).Methods(http.MethodPost)
	s.router.Path("/getPlayer").HandlerFunc(s.retrievePlayerInformation()).Methods(http.MethodPost)
	s.router.Path("/getAllPlayers").HandlerFunc(s.getAllPlayers()).Methods(http.MethodGet)
	s.router.Path("/deletePlayer").HandlerFunc(s.deletePlayer()).Methods(http.MethodPost)
}
