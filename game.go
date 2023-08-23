package main

import "github.com/gdamore/tcell/v2"

type Game struct {
	Screen tcell.Screen
	snakeBody SnakeBody
}