package packageone

import "fmt"

/*
var privateVariable = "I am private"
var PublicVariable = "I am public (or exported)"

func notExportedFunction() {
	fmt.Println("Not exported function")
}

func ExportedFunction() {
	fmt.Println("ExportedFunction function")
}
*/

var PackageVar = "package variable"

func PrintMe(s string) {
	fmt.Print(" (", s, ") ")
}
