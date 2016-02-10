package main

type Cell struct {
	x	int
	y	int
	Fg Attr // Foreground colour
	Bg Attr // Background color
	Ch rune // The character to draw
	content *Snake
}

func (c *Cell) glyph() rune {
	if c.content == nil {
		return ' '
	}
	
	return '*'
}
