package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
	FoodPos   Part
}

func drawParts(s tcell.Screen, parts []Part, foodPos Part, style tcell.Style) {
	s.SetContent(foodPos.X, foodPos.Y, '\u25CF', nil, style)
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

func checkCollision(parts []Part, otherPart Part) bool {
	for _, part := range parts {
		if part.X == otherPart.X && part.Y == otherPart.Y {
			return true
		}
	}
	return false
}

func (g *Game) Run() {
	width, height := g.Screen.Size()
	fmt.Println("width", width)
	fmt.Println("height", height)
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	g.snakeBody.ResetPos(width, height)
	g.UpdateFoodPos(width, height)
	for {
		longerSnake := false
		g.Screen.Clear()
		if checkCollision(g.snakeBody.Parts[len(g.snakeBody.Parts)-1:], g.FoodPos) {
			g.UpdateFoodPos(width, height)
			longerSnake = true
		}
		if checkCollision(g.snakeBody.Parts[:len(g.snakeBody.Parts)-1], g.snakeBody.Parts[len(g.snakeBody.Parts)-1]) {
			break
		}
		g.snakeBody.Update(width, height, longerSnake)
		drawParts(g.Screen, g.snakeBody.Parts, g.FoodPos, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
