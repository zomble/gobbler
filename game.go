package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

type Game struct {
	view		*View
	snakeHead	*Snake
	input		*Input
	grid		*Grid
}

func newGame(snake *Snake) *Game {

	grid := newGrid(10, 10)

	game := Game{
		snakeHead:	snake,
		view:	newView(grid),
		input:	newInput(),
		grid:	grid,
	}

	return &game
}

func (game *Game) tick(event Event) {
	// ???
}

func (game *Game) start() {

	err := termbox.Init()
//
//	termbox.SetOutputMode(termbox.Output256)
//	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

	if err != nil {
		panic(err)
	}

	defer termbox.Close()

//	game.screen.resize(termbox.Size())

	game.grid.getCenterCell().content = game.snakeHead

	game.input.start()
	defer game.input.stop()

	clock := time.Now()

	loop:
		for {
			tickTime := time.Now()
			game.view.delta = tickTime.Sub(clock).Seconds()
			clock = tickTime

			select {
				case event := <- game.input.queue:
					if event.Key == game.input.endKey {
						break loop
					} else if EventType(event.Type) == EventResize {
//						game.screen.resize(ev.Width, ev.Height)
					} else if EventType(event.Type) == EventError {
//						game.log(event.Err.Error())
						panic(event.Err)
					}
					game.view.tick(convertEvent(event))

					fmt.Printf("%+v\n", event)

				default:
					game.view.tick(Event{Type: EventNone})

					fmt.Println("nothing")
			}

			game.view.draw()

			time.Sleep(time.Duration((tickTime.Sub(time.Now()).Seconds()*1000.0)+1000.0/game.view.fps) * time.Millisecond)
		}
}
