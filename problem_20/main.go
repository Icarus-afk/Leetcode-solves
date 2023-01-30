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

func main() {
	var input string
	stack := &Stack{}
	fmt.Println("Enter values to add to the stack, separated by spaces:")
	for true {
		fmt.Scan(&input)
		if input == "\t" {
			break
		}
		stack.Push(input)

		fmt.Println("Stack:", stack.data)
	}
}
