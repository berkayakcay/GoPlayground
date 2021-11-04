package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚽'

		maxFrames = 1200
		speed     = time.Second / 30
	)

	var (
		px, py int
		vx, vy = 1, 1

		cell rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		board[px][py] = false
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		board[px][py] = true // set the ball position

		buf := make([]rune, 0, width*height)

		for y := range board[0] { // rows
			for x := range board { // columns
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}

			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))
		time.Sleep(speed)
	}

}
