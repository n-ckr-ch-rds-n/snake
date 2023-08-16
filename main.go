package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
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
var Screen tcell.Screen

var snake *Snake
var apple *Apple

const FRAME_WIDTH = 80
const FRAME_HEIGHT = 15
const SNAKE_SYMBOL = 0x2588
const APPLE_SYMBOL = 0x25CF

const FRAME_BORDER_THICKNESS = 1
const FRAME_BORDER_VERTICAL = '║'
const FRAME_BORDER_HORIZONTAL = '═'
const FRAME_BORDER_TOP_LEFT = '╔'
const FRAME_BORDER_TOP_RIGHT = '╗'
const FRAME_BORDER_BOTTOM_RIGHT = '╝'
const FRAME_BORDER_BOTTOM_LEFT = '╚'

func main() {
	initScreen()
	initGameObjects()
	displayFrame()
}

func initScreen() {
	var err error
	Screen, err = tcell.NewScreen()
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
	apple = &Apple{
		point:  getInitialAppleCoordinates(),
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

func displayFrame() {
	frameOriginX, frameOriginY := getFrameOrigin()
	printUnfilledRectangle(
		frameOriginX,
		frameOriginY,
		FRAME_WIDTH,
		FRAME_HEIGHT,
		FRAME_BORDER_THICKNESS,
		FRAME_BORDER_HORIZONTAL,
		FRAME_BORDER_VERTICAL,
		FRAME_BORDER_TOP_LEFT,
		FRAME_BORDER_TOP_RIGHT,
		FRAME_BORDER_BOTTOM_RIGHT,
		FRAME_BORDER_BOTTOM_LEFT,
	)
	Screen.Show()
}

func printUnfilledRectangle(
	xOrigin int,
	yOrigin int,
	width int,
	height int,
	borderThickness int,
	horizontalOutline rune,
	verticalOutline rune,
	topLeftOutline rune,
	topRightOutline rune,
	bottomRightOutline rune,
	bottomLeftOutline rune,
) {
	var upperBorder, lowerBorder rune
	verticalBorder := verticalOutline
	for i := 0; i < width; i++ {
		if i == 0 {
			upperBorder = topLeftOutline
			lowerBorder = bottomLeftOutline
		} else if i == width-1 {
			upperBorder = topRightOutline
			lowerBorder = bottomRightOutline
		} else {
			upperBorder = horizontalOutline
			lowerBorder = horizontalOutline
		}
		printToFrame(xOrigin+i, yOrigin, borderThickness, borderThickness, tcell.StyleDefault, upperBorder)
		printToFrame(xOrigin+i, yOrigin+height-1, borderThickness, borderThickness, tcell.StyleDefault, lowerBorder)
	}
	for i := 1; i < height-1; i++ {
		printToFrame(xOrigin, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
		printToFrame(xOrigin+width-1, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
	}
}

func printToFrame(x, y, w, h int, style tcell.Style, char rune) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			Screen.SetContent(x+i, y+j, char, nil, style)
		}
	}
}
