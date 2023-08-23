package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gdamore/tcell/v2"
)

type SnakeBody struct {
	X int
	Y int
	Xspeed int
	Yspeed int
}

func main() {
	fmt.Println("Loading snake ...")
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)
	snakeBody := SnakeBody{
		X: 5,
		Y: 10,
		Xspeed: 1,
		Yspeed: 0,
	}
}