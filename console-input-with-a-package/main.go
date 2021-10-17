package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {

	/*
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("->")

			userInput, _ := reader.ReadString('\n')

			userInput = strings.Replace(userInput, "\n", "", -1)

			if userInput == "quit" {
				break
			} else {
				fmt.Println(userInput)
			}
		}
	*/

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// as soon as main function ends it runs
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		// fmt.Sprintf("You chose %q", char)
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}

	fmt.Println("Program exiting..")
}
