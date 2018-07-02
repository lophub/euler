package main

import (
	"fmt"
)

func main() {
	x, y, sum := 1, 2, 2 //does shit
	for y < 4e6 {        // |
		x, y = y, x+y // |
		if y%2 == 0 { // |
			sum += y // v
		}
	}
	fmt.Println(sum)
}
