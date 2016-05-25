package main

import "github.com/nsf/termbox-go"

type View struct {
	grid      *Grid
	width     int
	height    int
	delta     float64
	fps       float64
}

func newView(g *Grid) *View {
	view := View{
		grid: 	g,
		fps:	30,
	}

	return &view
}

// Draw is called every frame by the game to render the current grid.
func (view *View) draw() {
	termbox.Clear(termbox.ColorBlue, termbox.ColorYellow)

//	view.futureFrame = newFrame(view.width, view.height)
//
//	for _, cell := range view.grid {
//		cell.draw(view)
//	}
//
//	// Check if anything changed between Draws
//	if !view.buffer.equals(&view.render) {
		// Draw to terminal
		for i, col := range view.grid.cells {
			for j, cell := range col {

				var glyph rune

				switch cell.s.(type) {
				default:
				case Placer:
					glyph = cell.s.(Placer).glyph()
				}

				termbox.SetCell(
					i,
					j,
					glyph,
					termbox.Attribute(cell.Fg),
					termbox.Attribute(cell.Bg),
				)
			}
		}

		termbox.Flush()
//	}

//	s.oldCanvas = s.canvas
}

//func (s *Screen) resize(w, h int) {
//	s.width = w
//	s.height = h
//	c := NewCanvas(w, h)
//	// Copy old data that fits
//	for i := 0; i < min(w, len(s.canvas)); i++ {
//		for j := 0; j < min(h, len(s.canvas[0])); j++ {
//			c[i][j] = s.canvas[i][j]
//		}
//	}
//	s.canvas = c
//}

// Size returns the width and height of the Screen
// in characters.
//func (s *Screen) Size() (int, int) {
//	return s.width, s.height
//}
//
//// TimeDelta returns the number of seconds since the previous
//// frame was rendered. Can be used for timings and animation.
//func (s *Screen) TimeDelta() float64 {
//	return s.delta
//}
//
//// Set the screen framerate.  By default, termloop will draw the
//// the screen as fast as possible, which may use a lot of system
//// resources.
//func (s *Screen) SetFps(f float64) {
//	s.fps = f
//}
//
//// RenderCell updates the Cell at a given position on the Screen
//// with the attributes in Cell c.
//func (s *Screen) RenderCell(x, y int, c *Cell) {
//	newx := x + s.offsetx
//	newy := y + s.offsety
//	if newx >= 0 && newx < len(s.canvas) &&
//	newy >= 0 && newy < len(s.canvas[0]) {
//		renderCell(&s.canvas[newx][newy], c)
//	}
//}
