package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	elements []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack) Pop() Flight {
	if (*s).IsEmpty() {
		return Flight{}
	}

	index := len((*s).elements) - 1
	element := (*s).elements[index]
	(*s).elements = (*s).elements[:len((*s).elements)-1]
	return element

}

func (s *Stack) Push(f Flight) {
	(*s).elements = append((*s).elements, f)
}

func (s *Stack) Peek() Flight {
	if (*s).IsEmpty() {
		return Flight{}
	}
	return (*s).elements[len((*s).elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len((*s).elements) == 0
}

func (s *Stack) String() string {
	stringBuilder := strings.Builder{}
	for _, value := range (*s).elements {
		stringBuilder.WriteString(fmt.Sprintf("%v\n", value))
	}
	return stringBuilder.String()
}

func main() {
	fmt.Println("Go Stack Implementation")

	stack := Stack{}
	stack.Push(Flight{
		Origin:      "IST",
		Destination: "UK",
		Price:       100,
	})

	fmt.Println(stack)

	stack.Push(Flight{
		Origin:      "IST",
		Destination: "GER",
		Price:       100,
	})

	stack.Push(Flight{
		Origin:      "IST",
		Destination: "BRA",
		Price:       100,
	})

	lastFlight := stack.Pop()
	fmt.Println(lastFlight)
	fmt.Println(stack)
}
