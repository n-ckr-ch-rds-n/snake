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

const FRAME_WIDTH = 80
const FRAME_HEIGHT = 15

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
