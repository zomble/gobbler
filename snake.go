package main

type SnakeState uint8

const (
	snakeHead SnakeState = iota
	snakeBody
)

type Snake struct {
	x         int
	y         int
	decay     int
	state     SnakeState
	direction Direction
}

func newSnake() *Snake {
	snake := Snake{
		decay: 1,
		state: snakeHead,
	}

	return &snake
}

func (snake *Snake) Ate() {
	snake.decay++
}

func (snake *Snake) Position() (int, int) {
	return snake.x, snake.y
}

func (snake *Snake) SetPosition(x, y int) {
	snake.x = x
	snake.y = y
}

func (s *Snake) glyph() rune {

	// TODO remove this.

	if s == nil {
		return ' '
	}

	if (s.state == snakeBody) {
		return '■'
	}

	switch s.direction {
	case directionDown:
		return '▼'
	case directionUp:
		return '▲'
	case directionLeft:
		return '◀'
	case directionRight:
		return '▶'
	}

	return '?'
}
