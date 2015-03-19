// Package gopicker is a simple winner picker
// created for a Codeaholics raffle
package gopicker

import (
	"fmt"
	"math/rand"
	"time"
)

// VERSION is the winner picker version
const VERSION = "0.0.2"

type competition struct {
	contestants []string
	rand        *rand.Rand
}

func newCompetition(contestants []string) *competition {
	return &competition{
		contestants: contestants,
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Perform presents the raffle, with a little pizzazz
func Perform(contestants []string) {
	competition := newCompetition(contestants)
	competition.perform()
}

func (c *competition) perform() {
	fmt.Println("WinnerPicker")
	fmt.Println("Version:", VERSION)
	fmt.Println("-----------------")
	fmt.Println("Contestants:")

	for _, name := range c.contestants {
		fmt.Println("*", name)
	}

	fmt.Println("-----------------")
	fmt.Println("Picking a winner:")
	fmt.Println("-----------------")

	for i := 10; i > 0; i-- {
		time.Sleep(time.Second)
		fmt.Println("#", i)
	}

	fmt.Println("-----------------")
	winner := PickWinner(c.contestants)
	fmt.Println("The winner is:")
	fmt.Println("*****************")
	fmt.Println(winner)
	fmt.Println("*****************")
}

// PickWinner takes a list of contestants, and picks a winner
func PickWinner(contestants []string) string {
	competition := newCompetition(contestants)
	return competition.pickWinner()
}

func (c *competition) pickWinner() string {
	if len(c.contestants) < 1 {
		panic("PickWinner needs >= 1 candidates")
	}

	winIndex := c.rand.Intn(len(c.contestants))
	return c.contestants[winIndex]
}
