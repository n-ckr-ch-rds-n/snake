package main

import "github.com/gdamore/tcell/v2/encoding"

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
	encoding.Register()
}
