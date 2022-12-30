package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	halfReverse := GetHalfReverse(&x)
	// 偶数 和 奇数的条件，符合一个就可以了
	return halfReverse == x || halfReverse/10 == x
}

func GetHalfReverse(x *int) int {
	result := 0
	for result < *x {
		lastN := *x % 10
		result = result*10 + lastN
		*x = *x / 10
	}
	return result
}
