package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	encountered := map[string]struct{}{}
	for _, item := range developers {
		encountered[item.Name] = struct{}{}
	}

	uniques := make([]string, 0, len(encountered))
	for key := range encountered {
		uniques = append(uniques, key)
	}
	return uniques
}

func main() {
	developers := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}
	uniques := FilterUnique(developers)
	fmt.Println(uniques)
	fmt.Println("Filter Unique Challenge")
}
