package main

import (
	"errors"
	"fmt"
)

/**
给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。



示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]
*/

func main() {
	temperatures := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := initStack(len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		temperature := temperatures[i]
		for stack.Top != -1 && stack.peek().value <= temperature {
			stack.pop()
		}
		if p := stack.peek(); p.value > temperature {
			result[i] = p.index - i
		} else {
			result[i] = 0
		}
		stack.push(&Node{
			index: i,
			value: temperature,
		})
	}

	return result
}

type Stack struct {
	MaxSize int
	Top     int
	arr     []*Node
}

type Node struct {
	value int
	index int
}

func initStack(size int) (stack *Stack) {
	stack = &Stack{
		MaxSize: size,
		Top:     -1,
		arr:     make([]*Node, size),
	}
	return stack
}

func (s *Stack) push(v *Node) error {
	//判断栈是否满
	if s.MaxSize-1 == s.Top {
		return errors.New("栈已满，无法插入数据")
	}

	s.Top++          //栈顶+1
	s.arr[s.Top] = v //入栈

	return nil
}

func (s *Stack) pop() (*Node, error) {
	//判断栈是否为空
	if s.Top == -1 {
		return &Node{
			index: 0,
			value: 0,
		}, errors.New("error: 栈为空")
	}

	v := s.arr[s.Top] //出栈
	s.Top--           //栈顶-1

	return v, nil
}

func (s *Stack) peek() *Node {
	//判断栈是否为空
	if s.Top == -1 {
		return &Node{
			index: 0,
			value: 0,
		}
	}

	v := s.arr[s.Top]

	return v
}
