package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		b = 1
		e = 1
		r = 1
		k = 1
		a = 1
		y = 1
		  = 1
		a = 1
		k = 1
		ç = 2 ***
		a = 1
		y = 1
	*/

	fmt.Println(len("berkay"))                    // 6
	fmt.Println(utf8.RuneCountInString("berkay")) // 6

	fmt.Println(len("akçay"))                    // 6
	fmt.Println(utf8.RuneCountInString("akçay")) // 5
}
