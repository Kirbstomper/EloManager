package main

import (
	"PlayerElo/elocaluclaton"
	"encoding/json"
	"log"
	"net/http"
)

var operator elocaluclaton.Operator

/*
	Adds players to the registry based on input given
*/

func AddPlayers(w http.ResponseWriter, r *http.Request) {
	var p []elocaluclaton.Player

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
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

func main() {
	operator = elocaluclaton.CreateInMemoryOperator()
	operator.CreateNewPlayer("chris", 1000)
	operator.CreateNewPlayer("paul", 1000)
	err := operator.RunMatch("paul", "chris", elocaluclaton.WIN)
	if err != nil {
		log.Println(err.Error())
	}
	operator.RetrievePlayerInformation("paul")
	operator.RetrievePlayerInformation("chris")

	http.HandleFunc("/add", AddPlayers)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
