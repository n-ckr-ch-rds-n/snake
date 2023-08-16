package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Coordinate struct {
	x, y int
}

type Snake struct {
	points                      []*Coordinate
	columnVelocity, rowVelocity int
	symbol                      rune
}

type Apple struct {
	point  *Coordinate
	symbol rune
}

var screenWidth, screenHeight int

const FRAME_WIDTH = 80
const FRAME_HEIGHT = 15
const FRAME_BORDER_THICKNESS = 1
const SNAKE_SYMBOL = 0x2588

func main() {
	initScreen()
}

func initScreen() {
	var err error
	Screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = Screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	Screen.SetStyle(defStyle)
	screenWidth, screenHeight := Screen.Size()
	if screenWidth < FRAME_WIDTH || screenHeight < FRAME_HEIGHT {
		fmt.Printf("The game frame is defined with %d width and %d height. Increase terminal size.", FRAME_WIDTH, FRAME_HEIGHT)
		os.Exit(1)
	}
}

func initGameObjects() {
	snake = &Snake{
		points:         getInitialSnakeCoordinates(),
		columnVelocity: 0,
		rowVelocity:    1,
		symbol:         SNAKE_SYMBOL,
	}
}

func getInitialSnakeCoordinates() []*Coordinate {
	snakeInitialCoordinate1 := &Coordinate{8, 4}
	transformInitialCoordinateInsideFrame(snakeInitialCoordinate1)
}

func transformInitialCoordinateInsideFrame(c *Coordinate) {
	leftX, topY, rightX, bottomY := getBoundaries()

}

func getBoundaries() (int, int, int, int) {
	originX, originY := getFrameOrigin()
	topY := originY
	bottomY := originY + FRAME_HEIGHT - FRAME_BORDER_THICKNESS
	leftX := originX
	rightX := originX + FRAME_WIDTH - FRAME_BORDER_THICKNESS
	return leftX, topY, rightX, bottomY
}

func getFrameOrigin() (int, int) {
	return (screenWidth-FRAME_WIDTH)/2 - FRAME_BORDER_THICKNESS, (screenHeight-FRAME_HEIGHT)/2 - FRAME_BORDER_THICKNESS
}
