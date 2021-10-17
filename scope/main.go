package main

import (
	"GoPlayground/scope/packageone"
	"fmt"
)

var myVariable = "my variable"

func main() {
	// variables
	// declare a package level variable for the main
	// package named myVar

	// declare a block variable for the main function
	// called blockVar
	var blockVariable = "block variable"

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in packageone called PrintMe
	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on one line using the PrintMe
	// function in packageone

	fmt.Println("Print start..")
	packageone.PrintMe(myVariable)
	packageone.PrintMe(blockVariable)
	packageone.PrintMe(packageone.PackageVar)
	fmt.Println("")
	fmt.Println("Print end..")

}
