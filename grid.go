package main

import "math"

type GridCells [][]Cell

type Grid struct {
	width	int
	height  int
	cells	GridCells
}

func newGrid(width, height int) *Grid {

	cells := make(GridCells, width)
	for i := range cells {
		cells[i] = make([]Cell, height)
	}

	grid := Grid{
		width:	width,
		height:	height,
		cells:	cells,
	}

	return &grid
}

func (g *Grid) decay() {

}

func (g *Grid) getCenterCell() *Cell {
	x := int(math.Floor(float64(g.width / 2)))
	y := int(math.Floor(float64(g.height / 2)))

	return &g.cells[x][y]
}
