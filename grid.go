package main

import (
	"math"
)

type Positioner interface {
	Position() (int, int)
	SetPosition(x, y int)
}

type Direction uint8

const (
	directionUp    Direction = iota
	directionDown
	directionLeft
	directionRight
)

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

func (g *Grid) Cell(x, y int) *Cell {
	if (0 > y || 0 > x || x > len(g.cells)-1 || y > len(g.cells[x])-1) {
		return nil
	}

	return &g.cells[x][y]
}

func (g *Grid) CenterPosition() (int, int) {
	return int(math.Floor(float64(g.width / 2))), int(math.Floor(float64(g.height / 2)))
}

func (g *Grid) PositionAdjacent(x int, y int, d Direction) (int, int)  {
	switch d {
	case directionUp:
		y = y-1
	case directionDown:
		y = y+1
	case directionLeft:
		x = x-1
	case directionRight:
		x = x+1
	}

	return x, y
}
