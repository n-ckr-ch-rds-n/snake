package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

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

	snakeParts := []Part{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
	}

	snakeBody := SnakeBody{
		Parts:  snakeParts,
		Xspeed: 1,
		Yspeed: 0,
	}

	game := Game{
		Screen:    screen,
		snakeBody: snakeBody,
		Score:     0,
		GameOver:  false,
	}
	go game.Run()
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.snakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				game.snakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.snakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.snakeBody.ChangeDir(0, 1)
			} else if event.Rune() == 'y' && game.GameOver {
				go game.Run()
			} else if event.Rune() == 'n' && game.GameOver {
				game.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
