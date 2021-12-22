package main

import (
	"PlayerElo/elocaluclaton"
	"encoding/json"
	"log"
	"net/http"
)

type Match struct {
	PlayerA, PlayerB, Result string
}

var operator elocaluclaton.Operator

/*
	Adds one or more players to the registry based on input given
*/

func AddPlayers(w http.ResponseWriter, r *http.Request) {
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
		err := operator.CreateNewPlayer(x.Tag, x.Elo)
		if err != nil {
			errors += err.Error() + "\n"
			hasErrors = true
		}

	}
	if hasErrors {
		http.Error(w, errors, http.StatusBadRequest)
	}
}

func Decide(w http.ResponseWriter, r *http.Request) {
	var m Match

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = operator.RunMatch(m.PlayerA, m.PlayerB, elocaluclaton.ResultsMap[m.Result])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdatePlayerElo(w http.ResponseWriter, r *http.Request) {
	var p elocaluclaton.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = operator.UpdateEloForPlayer(p.Tag, p.Elo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func RetrievePlayerInformation(w http.ResponseWriter, r *http.Request) {
	var p elocaluclaton.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p, err = operator.RetrievePlayerInformation(p.Tag)
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

func main() {

	operator = elocaluclaton.CreateRealtionalDBOperator()
	operator.CreateNewPlayer("chris", 1000)
	operator.CreateNewPlayer("paul", 1000)
	operator.RetrievePlayerInformation("paul")
	operator.RetrievePlayerInformation("chris")

	http.HandleFunc("/add", AddPlayers)
	http.HandleFunc("/decide", Decide)
	http.HandleFunc("/updateElo", UpdatePlayerElo)
	http.HandleFunc("/getPlayer", RetrievePlayerInformation)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
