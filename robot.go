package main

type Robot struct {
    posX, posY int
    orientation string
}

// A Robot must have awareness of the World he's in.
// Physical analogy would be, e.g., a paper map.
func NewRobot(x int, y int, o string, w World) Robot {
    temp := Robot{posX: x, posY: y, orientation: o, world: w}
    return temp
}

func (r *Robot) Forward {
}

func (r *Robot) TurnLeft() {
}

func (r *Robot) TurnRight() {
}
