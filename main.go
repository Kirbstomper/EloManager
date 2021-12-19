package main

import calc "PlayerElo/elocaluclaton"

func main() {
	p1 := 1000
	p2 := 1100
	println("Player 1: ", p1, " | Player 2 ", p2)

	p1, p2 = calc.CalculateElo(p1, p2, calc.TIE)
	println("Player 1: ", p1, " | Player 2 ", p2)

	p1, p2 = calc.CalculateElo(p1, p2, calc.TIE)
	println("Player 1: ", p1, " | Player 2 ", p2)

	p1, p2 = calc.CalculateElo(p1, p2, calc.LOSS)
	println("Player 1: ", p1, " | Player 2 ", p2)

}
