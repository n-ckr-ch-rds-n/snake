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
const APPLE_SYMBOL = 0x25CF

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
	apple := &Apple{
		point: getInitialAppleCoordinates(),
		symbol: APPLE_SYMBOL,
	}
}

func getInitialSnakeCoordinates() []*Coordinate {
	snakeInitialCoordinate1 := &Coordinate{8, 4}
	transformCoordinateInsideFrame(snakeInitialCoordinate1)
	snakeInitialCoordinate2 := &Coordinate{8, 5}
	transformCoordinateInsideFrame(snakeInitialCoordinate2)
	snakeInitialCoordinate3 := &Coordinate{8, 6}
	transformCoordinateInsideFrame(snakeInitialCoordinate3)
	snakeInitialCoordinate4 := &Coordinate{8, 7}
	transformCoordinateInsideFrame(snakeInitialCoordinate4)
	return []*Coordinate{
		{snakeInitialCoordinate1.x, snakeInitialCoordinate1.y},
		{snakeInitialCoordinate2.x, snakeInitialCoordinate2.y},
		{snakeInitialCoordinate3.x, snakeInitialCoordinate3.y},
		{snakeInitialCoordinate4.x, snakeInitialCoordinate4.y},
	}
}

func getInitialAppleCoordinates() *Coordinate {
	appleInitialCoordinate := &Coordinate{FRAME_WIDTH / 2, FRAME_HEIGHT / 2}
	transformCoordinateInsideFrame(appleInitialCoordinate)
	return appleInitialCoordinate
}

func transformCoordinateInsideFrame(c *Coordinate) {
	leftX, topY, rightX, bottomY := getBoundaries()
	c.x += leftX + FRAME_BORDER_THICKNESS
	c.y += topY + FRAME_BORDER_THICKNESS
	for c.x >= rightX {
		c.x--
	}
	for c.y >= bottomY {
		c.y--
	}

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
