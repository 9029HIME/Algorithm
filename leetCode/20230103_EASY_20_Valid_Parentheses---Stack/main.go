package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("){"))
}

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	stack := InitStack()
	for _, v := range s {
		p := string(v)
		// 左括号，直接压栈
		if p == "{" || p == "(" || p == "[" {
			stack.push(p)
			continue
		}
		// 像}(这种结构，一开始就给右括号，默认pop返回error，否则-1下标会抛异常。
		pop := stack.pop()
		if pop == "error" {
			return false
		}

		// 比较右括号 和 左括号是否匹配，有一个不匹配就返回false
		switch p {
		case "}":
			if pop != "{" {
				return false
			}
		case ")":
			if pop != "(" {
				return false
			}
		case "]":
			if pop != "[" {
				return false
			}
		default:
			return false
		}
	}
	// 有可能s全部由左括号构成，此时stack的数据长度不为空，返回false
	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}

type Stack struct {
	data []string
}

func InitStack() *Stack {
	return &Stack{
		data: make([]string, 0, 8),
	}
}

func (this *Stack) pop() string {
	index := len(this.data) - 1
	if index == -1 {
		return "error"
	}
	pop := this.data[index]
	this.data = this.data[0:index]
	return pop
}

func (this *Stack) push(str string) {
	this.data = append(this.data, str)
}

func (this *Stack) isEmpty() bool {
	return len(this.data) == 0
}
