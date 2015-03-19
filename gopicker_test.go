package gopicker

import (
	"testing"

	"github.com/matthewrudy/set"
)

func Test_PickWinner(t *testing.T) {
	candidates := []string{
		"Alice",
		"Bob",
		"Charlie",
	}

	winners := set.New()
	iterations := len(candidates) * 10

	for i := 0; i < iterations; i++ {
		winners.Add(PickWinner(candidates))
	}

	if winners.Size() != len(candidates) {
		t.Errorf("after %v iterations, each candidate should be chosen: %v", iterations, winners)
	}
}

func Test_PickWinner_oneChoice(t *testing.T) {
	candidates := []string{
		"Alice",
	}

	winner := PickWinner(candidates)

	if winner != "Alice" {
		t.Error("Winner should be Alice")
	}
}

// func Test_PickWinner_noChoice(t *testing.T) {
// 	candidates := []string{}
//
// 	PickWinner(candidates)
// }
