package main

import "github.com/matthewrudy/gopicker"

func main() {
	CONTESTANTS := []string{
		"Eddie",
		"Steve",
		"Matthew",
	}

	gopicker.Pick(CONTESTANTS)
}
