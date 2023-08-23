package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
}

func (g *Game) Run() {
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		g.Screen.Clear()
		g.snakeBody.Update(width, height)
		// draw parts
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
