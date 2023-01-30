package main

import (
	"fmt"
)

type Stack struct {
	data []string
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() string {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Peek() string {
	return s.data[len(s.data)-1]
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(s.data) - 1
		elem := (s.data)[index]
		return elem
	}
}

func isValid(s *Stack) bool {
	stack := new(Stack)
	x := len(stack.data)
	for v := range x {
		if v == 40 || v == 91 || v == 123 {
			stack.Push(string(v))
		}

		if v == 41 {
			if stack.Top() == "(" {
				stack.Pop()
			} else {
				return false
			}
		}
		if v == 93 {
			if stack.Top() == "[" {
				stack.Pop()
			} else {
				return false
			}
		}
		if v == 125 {
			if stack.Top() == "{" {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	if len := stack.Len(); len > 0 {
		return false
	}
	return true
}

func main() {
	var input string
	stack := &Stack{}
	fmt.Println("Enter values to add to the stack, separated by spaces:")
	for true {
		fmt.Scan(&input)
		if input == "~" { //temporary solution for breaking out of the loop
			break
		}
		stack.Push(input)

		fmt.Println("Stack:", stack.data)
	}
	fmt.Println(isValid(stack))
}
