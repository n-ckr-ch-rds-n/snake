package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
	FoodPos   Part
}

func drawParts(s tcell.Screen, parts []Part, style tcell.Style) {
	for _, part := range parts {
		s.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func (g *Game) UpdateFoodPos(width, height int) {
	g.FoodPos.X = rand.Intn(width)
	g.FoodPos.Y = rand.Intn(height)
	if g.FoodPos.Y == 1 && g.FoodPos.X < 10 {
		g.UpdateFoodPos(width, height)
	}
}

func (g *Game) Run() {
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	g.snakeBody.ResetPos(width, height)
	g.UpdateFoodPos(width, height)
	for {
		longerSnake := false
		g.Screen.Clear()
		g.snakeBody.Update(width, height, longerSnake)
		drawParts(g.Screen, g.snakeBody.Parts, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
