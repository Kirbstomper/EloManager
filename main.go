package main

import "PlayerElo/elocaluclaton"

var operator elocaluclaton.Operator

func main() {
	operator = elocaluclaton.CreateInMemoryOperator()
	operator.CreateNewPlayer("chris", 1000)
	operator.CreateNewPlayer("paul", 1000)

	print(operator.CalculateEloForPlayers("paul", "chris", elocaluclaton.WIN))
}
