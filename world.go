package main

const MaxX = 50
const MaxY = 50

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
    temp := World{BoundX: x, BoundY: y}
    temp.Grid = make([]Cells, MaxY)
    for i:=0; i<MaxY; i++ {
        temp.Grid[i] = make(Cells, MaxX)
    }
    return temp
}

func (w World) HasScent(x, y int) bool {
    return w.Grid[y][x].Scent
}

func (w *World) AddScent(x, y int) {
    w.Grid[y][x].Scent = true
}
