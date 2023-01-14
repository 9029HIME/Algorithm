package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses("(()") == 2)
	fmt.Println(longestValidParentheses(")()())") == 4)
	fmt.Println(longestValidParentheses("") == 0)
	fmt.Println(longestValidParentheses("()(()") == 2)
	fmt.Println(longestValidParentheses("()(())") == 6)

}

func longestValidParentheses(s string) int {
	stack := InitStack()
	result := 0
	tmp := 0
	stack.push(-1)
	for i, c := range s {
		s1 := string(c)
		if s1 == "(" {
			stack.push(i)
			continue
		}
		if s1 == ")" {
			tailIndex := stack.tail()
			if tailIndex != -1 && string(s[tailIndex]) == "(" {
				// 先pop，再减
				stack.pop()
				tmp = i - stack.tail()
				result = Max(tmp, result)
			} else {
				stack.push(i)
			}
		}

	}
	result = Max(tmp, result)
	return result
}

type Stack struct {
	data []int
}

func InitStack() *Stack {
	return &Stack{
		data: make([]int, 0, 8),
	}
}

func (this *Stack) pop() int {
	index := len(this.data) - 1
	if index == -1 {
		return -1
	}
	pop := this.data[index]
	this.data = this.data[0:index]
	return pop
}

func (this *Stack) tail() int {
	index := len(this.data) - 1
	if index == -1 {
		return -1
	}
	pop := this.data[index]
	return pop
}

func (this *Stack) push(str int) {
	this.data = append(this.data, str)
}

func (this *Stack) isEmpty() bool {
	return len(this.data) == 0
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
