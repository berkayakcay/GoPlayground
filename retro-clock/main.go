package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	type placeHolder [5]string

	zero := placeHolder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeHolder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeHolder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeHolder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeHolder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeHolder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeHolder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeHolder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeHolder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeHolder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeHolder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeHolder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]placeHolder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
		Clear()
	}

}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Clear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
