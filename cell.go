package main

type Placer interface {
	glyph() rune // ???
}

type Cell struct {
	Fg uint16 // Foreground colour
	Bg uint16 // Background color
	Ch rune // The character to draw
	s  interface{}
}
