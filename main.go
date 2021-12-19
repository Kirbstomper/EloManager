package main

import (
	"PlayerElo/elocaluclaton"
)

var operator elocaluclaton.Operator

func main() {
	operator = elocaluclaton.CreateInMemoryOperator()
	operator.CreateNewPlayer("chris", 1000)
	operator.CreateNewPlayer("paul", 1000)
	err := operator.RunMatch("paul", "chris", elocaluclaton.WIN)
	if err != nil {
		panic(err.Error())
	}
	operator.RetrievePlayerInformation("paul")
	operator.RetrievePlayerInformation("chris")
}
