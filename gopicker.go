// Package gopicker is a simple winner picker
// created for a Codeaholics raffle
package gopicker

import (
	"fmt"
	"math/rand"
	"time"
)

// VERSION is the winner picker version
const VERSION = "0.0.1"

// Show takes a list of contestants, and creates a show to STDOUT
func Show(contestants []string) {
	fmt.Println("WinnerPicker")
	fmt.Println("Version:", VERSION)
	fmt.Println("-----------------")
	fmt.Println("Contestants:")

	for _, name := range contestants {
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
	winner := PickWinner(contestants)
	fmt.Println("The winner is:")
	fmt.Println("*****************")
	fmt.Println(winner)
	fmt.Println("*****************")
}

// PickWinner takes a list of contestants, and picks a winner
func PickWinner(contestants []string) string {
	rand.Seed(time.Now().Unix())
	winIndex := rand.Intn(len(contestants))
	return contestants[winIndex]
}
