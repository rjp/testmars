package main

import (
	"fmt"
)

// Define types for `Cell` and `Cells` because not only
// does it help with type safety, it makes the code easier
// to read for someone else.
type Cell struct {
	Scent bool
}

type Cells []Cell

type World struct {
	BoundX, BoundY int
	// An array of array of cells => a 2D array of cell
	Grid []Cells
}

func NewWorld(x, y int) World {
	// Do we need to check whether we're creating an axis bigger than 50?
	temp := World{BoundX: x, BoundY: y}
	// Since `y` is the last cell, we need to make `y+1` to account for 0-based.
	temp.Grid = make([]Cells, y+1)
	for i := 0; i < y+1; i++ {
		temp.Grid[i] = make(Cells, x+1)
	}
	return temp
}

func NewWorldFromString(s string) World {
	var x, y int
	// This is slightly flawed since you can feed it '3 4 5' and
	// go will happily parse out the '3 4' and ignore the 5. Could
	// be worked around with `%d %d%s`, maybe, and checking if
	// there was anything in the `%s`? Or splitting the line into
	// `Fields`, checking the number, and `ParseInt`ing each one.
	n, err := fmt.Sscanf(s, "%d %d", &x, &y)
	if err != nil || n != 2 {
		panic("Parsing world definition string")
	}
	return NewWorld(x, y)
}

// Helper functions because we don't want robots to have
// direct access to the grid just in case we want to mock
// or replace these with db access, etc.
func (w World) HasScent(x, y int) bool {
	return w.Grid[y][x].Scent
}

func (w *World) AddScent(x, y int) {
	w.Grid[y][x].Scent = true
}
