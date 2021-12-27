package handlers

import (
	"PlayerElo/elocaluclaton"
	"encoding/json"
	"net/http"
)

type Match struct {
	PlayerA, PlayerB, Result string
}

/*
	Adds one or more players to the registry based on input given
*/

func (s defaultServer) addPlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p []elocaluclaton.Player

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}
		println("Players retrived ", len(p))
		errors := ""
		hasErrors := false
		for _, x := range p {
			err := s.op.CreateNewPlayer(x.Tag, x.Elo)
			if err != nil {
				errors += err.Error() + "\n"
				hasErrors = true
			}

		}
		if hasErrors {
			http.Error(w, errors, http.StatusBadRequest)
		}
	}
}

func (s defaultServer) decide() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Match

		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.op.RunMatch(m.PlayerA, m.PlayerB, elocaluclaton.ResultsMap[m.Result])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s defaultServer) updatePlayerElo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p elocaluclaton.Player
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.op.UpdateEloForPlayer(p.Tag, p.Elo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s defaultServer) retrievePlayerInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p elocaluclaton.Player
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p, err = s.op.RetrievePlayerInformation(p.Tag)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
