package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	userName := readString("What is your name?")
	age := readInt("How old are you?")
	fmt.Println("Your name is", userName, "and you are", age, "years old.")
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func readInt(s string) int {
	fmt.Println(s)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Please enter a whole number")
	}

	return num
}
