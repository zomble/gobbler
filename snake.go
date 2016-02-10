package main

type SnakeDirection uint8

type Snake struct {
	decay		int
	direction 	SnakeDirection
}

func newSnake() *Snake {
	snake := Snake{
		decay: 1,
		direction: 1,
	}

	return &snake
}

func (snake *Snake) ate() {
	snake.decay++
}
