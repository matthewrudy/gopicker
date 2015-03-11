package main

import "github.com/matthewrudy/gopicker"

func main() {
	CONTESTANTS := []string{
		"Bart",
		"Jun Jun",
		"Zak",
		"Thomas",
		"Walter",
		"Jack",
		"Bambi",
		"Doug",
	}

	gopicker.Perform(CONTESTANTS)
}
