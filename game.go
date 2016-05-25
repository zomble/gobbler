package main

import (
	"time"
	"fmt"
)

type Game struct {
	snakeHead	*Snake
	grid		*Grid
	moveTimer   *time.Timer
}

func newGame(grid *Grid, snake *Snake) *Game {
	g := Game{
		snakeHead:	snake,
		grid: grid,
		moveTimer: time.NewTimer(time.Duration(time.Second)),
	}

	g.moveTimer.Stop()

	return &g
}

func (g *Game) move() {
	// reset the timer
	g.resetMoveTimer()

	x, y := g.snakeHead.Position();

	// leave the body behind
	body := newSnake()
	body.state = snakeBody
	body.decay = g.snakeHead.decay
	body.SetPosition(x, y)
	g.grid.Cell(x, y).s = body

	// ???
	xNew, yNew := g.grid.PositionAdjacent(x, y, g.snakeHead.direction)
	nextCell := g.grid.Cell(xNew, yNew)

	if (nextCell == nil) {
		fmt.Println("WTF")
	}

//	switch nextCell.(type) {
//	case nil:
		g.snakeHead.SetPosition(xNew, yNew)
		nextCell.s = g.snakeHead
//	default:
//		fmt.Println("WTF")
//	}
}

func (g *Game) start() {
	// start the timers and shit
	g.resetMoveTimer()
}

func (g *Game) resetMoveTimer() {
	g.moveTimer.Reset(time.Duration(time.Second))
}
