package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	length := len(nums)
	for slow < length && fast < length {
		v := nums[fast]
		if v != val {
			nums[slow] = v
			slow++
		}
		fast++
	}
	return slow
}
