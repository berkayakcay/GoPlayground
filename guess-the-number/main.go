package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and ENTER press when ready."

func main() {

	// Seed the number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {

	reader := bufio.NewReader(os.Stdin)

	// Display a welcome / instructions
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	_, _ = reader.ReadString('\n')

	// Take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	_, _ = reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	_, _ = reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer, prompt)
}
