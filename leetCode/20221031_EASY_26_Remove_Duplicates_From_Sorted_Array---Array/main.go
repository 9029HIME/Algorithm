package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	slow := 0
	fast := 0

	for fast < length {

		sE := nums[slow]
		fE := nums[fast]

		if sE != fE {
			slow++
			nums[slow] = fE
		}
		fast++
	}

	return slow + 1
}
