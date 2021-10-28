package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	/*
		Rather than using variable use struct example
	*/
	//userName := readString("What is your name?")
	//age := readInt("How old are you?")
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	// Easy but not best
	//fmt.Println("Your name is", userName, "and you are", age, "years old.")

	// You can not concatenate two different type (string, int)
	//fmt.Println("Your name is " +  userName + " and you are " + age + " years old.")

	// More efficient way
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, age))
	// More resource efficient
	fmt.Printf(
		"Your name is %s. You are %d years old. Your favourite number is %.2f. Owns a dog %t \n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnsADog,
	)

	/*
		GoLang fmt Printing Cheat Sheet by gpascual
		https://cheatography.com/gpascual/cheat-sheets/golang-fmt-printing/
	*/
}

func prompt() {
	fmt.Print("->")
}

func getUserInput() string {
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {

	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		lowerChar := strings.ToLower(string(char))
		if lowerChar != "y" && lowerChar != "n" {
			fmt.Println("Please enter y or n")
		} else if lowerChar == "n" {
			return false
		} else if lowerChar == "y" {
			return true
		}
	}
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput := getUserInput()
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput := getUserInput()

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {

			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput := getUserInput()

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {

			return num
		}
	}
}
