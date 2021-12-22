package elocaluclaton

import "math"

/*
	Handles all the elo calculaton when passed two score, and results
*/

const (
	LOSS = iota - 1
	WIN  = iota
	TIE  = iota

	K = 32
)

var ResultsMap = map[string]int{"WIN": WIN, "LOSS": LOSS, "TIE": TIE}

/*
	Calculates and returns new elo ratings based on the outcome
	p1: player 1's elo rating
	p2: player 2's elo rating
	outcome: the outcome for player 1
*/
func CalculateElo(pa int, pb int, outcome int) (int, int) {

	expectedA, expectedB := CalculateExpected(pa, pb), CalculateExpected(pb, pa)

	switch outcome {
	case LOSS:
		return CalculateNewRating(pa, K, 0, expectedA), CalculateNewRating(pb, K, 1, expectedB)
	case WIN:
		return CalculateNewRating(pa, K, 1, expectedA), CalculateNewRating(pb, K, 0, expectedB)
	case TIE:
		return CalculateNewRating(pa, K, 0.5, expectedA), CalculateNewRating(pb, K, 0.5, expectedB)
	}
	return -1, -1
}

/**
Calculate the expected score for PA given 2 elo ratings
*/
func CalculateExpected(pa int, pb int) float64 {
	expected := 1 / (1 + math.Pow(10, (float64(pb)-float64(pa))/400))

	return expected
}

/*
Calculates the new rating for a player
*/
func CalculateNewRating(current int, k int, score float64, expected float64) int {

	newRating := float64(current) + float64(k)*(score-expected)
	return int(newRating)
}

