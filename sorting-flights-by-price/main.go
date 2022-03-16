package main

import (
	"fmt"
	sort "sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type Programmer struct {
	Age int
}

type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	// implement
	sort.Slice(
		flights,
		func(i, j int) bool {
			fmt.Printf("i: %d, j: %d, return: %v \n", flights[i].Price, flights[j].Price, flights[i].Price < flights[j].Price)
			return flights[i].Price > flights[j].Price
		},
	)

	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d \n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {

	programmers := []Programmer{
		{Age: 1000},
		{Age: 50},
		{Age: 30},
		{Age: 20},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)

	// an empty slice of flights
	var flights []Flight
	flights = append(flights, Flight{Price: 10})
	flights = append(flights, Flight{Price: 50})
	flights = append(flights, Flight{Price: 30})

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
