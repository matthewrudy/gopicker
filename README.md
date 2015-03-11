gopicker
========

Pick a Competition Winner... with a little style

``` go
package main

import "github.com/matthewrudy/gopicker"

func main() {
	CONTESTANTS := []string{
		"Rod",
		"Jane",
		"Freddy",
	}

	gopicker.Perform(CONTESTANTS)
}
```

Run it;

``` bash
WinnerPicker
Version: 0.0.2
-----------------
Contestants:
* Rod
* Jane
* Freddy
-----------------
Picking a winner:
-----------------
# 10
# 9
# 8
# 7
# 6
# 5
# 4
# 3
# 2
# 1
-----------------
The winner is:
*****************
Jane
*****************
```

BOOM!
