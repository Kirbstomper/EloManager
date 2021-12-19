package elocaluclaton

import (
	"testing"
)

func TestCalculateWhenWin(t *testing.T) {
	a, b := CalculateElo(100, 100, WIN)
	if a <= 100 {
		t.Fail()
	}
	if b >= 100 {
		t.Fail()
	}
}

func TestCalculateWhenLoss(t *testing.T) {
	a, b := CalculateElo(100, 100, LOSS)
	if a >= 100 {
		t.Fail()
	}
	if b <= 100 {
		t.Fail()
	}
}

func TestCalculateWhenAdvantageTie(t *testing.T) {
	a, b := CalculateElo(200, 100, TIE)
	if a >= 200 {
		t.Fail()
	}
	if b <= 100 {
		t.Fail()
	}
}

func TestCalculateWhenDisadvantageTie(t *testing.T) {
	a, b := CalculateElo(100, 200, TIE)
	if a <= 100 {
		t.Fail()
	}
	if b >= 200 {
		t.Fail()
	}
}
func TestCalculateExpectedWhenEqual(t *testing.T) {
	expected := CalculateExpected(1000, 1000)
	if expected != .5 {
		t.Fail()
	}
}

func TestCalculateExpectedWhenOneHigher(t *testing.T) {
	expectedA := CalculateExpected(1000, 1500)
	expectedB := CalculateExpected(1500, 1000)

	if expectedA >= expectedB {
		t.Fail()
	}
}

func TestCalculateNewRatingWin(t *testing.T) {
	current := 1000
	k := 32
	score := 1.0
	expected := .6

	new := CalculateNewRating(current, k, score, expected)

	if new <= current {
		t.Fail()
	}

}

func TestCalculateNewRatingLoss(t *testing.T) {
	current := 1000
	k := 32
	score := 0.0
	expected := .6

	new := CalculateNewRating(current, k, score, expected)

	if new >= current {
		t.Fail()
	}

}
