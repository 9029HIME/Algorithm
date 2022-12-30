package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-123456))
}

func reverse(x int) int {
	negative := false
	if x < 0 {
		x = -x
		negative = true
	}

	result := 0
	for x != 0 {
		lastN := x % 10
		result = result * 10
		result = result + lastN
		if result > math.MaxInt32 {
			return 0
		}
		x = x / 10
	}
	if negative {
		result = -result
	}
	return result
}
