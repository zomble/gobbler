package main

func main() {
	snakeHead := newSnake()

	game := newGame(snakeHead)

	game.start()
}
