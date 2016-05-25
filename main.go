package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	// set up the game
	grid := newGrid(30, 15)
	input := newInput()
	view := newView(grid)
	snakeHead := newSnake()
	snakeHead.SetPosition(grid.CenterPosition())
	grid.Cell(grid.CenterPosition()).s = snakeHead

	game := newGame(grid, snakeHead)

	err := termbox.Init()
	//	termbox.SetOutputMode(termbox.Output256)
	//	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	input.start()
	defer input.stop()

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case ev := <- input.queue:
				if ev.Type == termbox.EventKey {
					switch ev.Key {
					case termbox.KeyArrowLeft:
						snakeHead.direction = directionLeft
					case termbox.KeyArrowRight:
						snakeHead.direction = directionRight
					case termbox.KeyArrowUp:
						snakeHead.direction = directionUp
					case termbox.KeyArrowDown:
						snakeHead.direction = directionDown
					case termbox.KeySpace:
						game.move()
					case termbox.KeyEsc, termbox.KeyCtrlC, termbox.KeyCtrlD:
						close(quit)
						return
					default:
						switch ev.Ch {
							case 'q':
								close(quit)
								return
							case 's':
								game.start()
//							case 'p':
//								game.pause()
						}
					}
				}
			case <- game.moveTimer.C:
				game.move()
			}
		}
	}()

	clock := time.Now()
	loop:
	for {
		select {
		default:
			tickTime := time.Now()
			view.delta = tickTime.Sub(clock).Seconds()
			clock = tickTime

			view.draw()

			time.Sleep(time.Duration((tickTime.Sub(time.Now()).Seconds()*1000.0)+1000.0/view.fps) * time.Millisecond)
		case <- quit:
			break loop;
		}
	}
}
