package main

import (
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
	point *Coordinate
	symbol rune
}

func main() {

}

func initScreen() {
	var err error
	Screen, err := tcell.NewScreen()
}
