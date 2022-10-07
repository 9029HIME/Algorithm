package main

import (
	"errors"
)

func main() {
}

type Stack struct {
	MaxSize int
	Top     int
	arr     [10]int
}

func initStack() (stack *Stack) {
	stack = &Stack{
		MaxSize: 10,
		Top:     -1,
		arr:     [10]int{},
	}
	return stack
}

func (s *Stack) push(v int) error {
	//判断栈是否满
	if s.MaxSize-1 == s.Top {
		return errors.New("栈已满，无法插入数据")
	}

	s.Top++          //栈顶+1
	s.arr[s.Top] = v //入栈

	return nil
}

func (s *Stack) pop() (int, error) {
	//判断栈是否为空
	if s.Top == -1 {
		return 0, errors.New("error: 栈为空")
	}

	v := s.arr[s.Top] //出栈
	s.Top--           //栈顶-1

	return v, nil
}

func (s *Stack) peek() (int, error) {
	//判断栈是否为空
	if s.Top == -1 {
		return 0, errors.New("error: 栈为空")
	}

	v := s.arr[s.Top]

	return v, nil
}
