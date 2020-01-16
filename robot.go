package main

type Robot struct {
    posX, posY int
    orientation string
    world World
}

// A Robot must have awareness of the World he's in.
// Physical analogy would be, e.g., a paper map.
func NewRobot(x int, y int, o string, w World) Robot {
    temp := Robot{posX: x, posY: y, orientation: o, world: w}
    return temp
}

func (r Robot) Position() (x, y int) {
    return r.posX, r.posY
}

// Translate a given orientation into specific axis changes.
// In a more complicated scenario, we might model this with
// a rotational matrix and combine that with a unit matrix to
// calculate our new position. But this is a simple case and
// we'll just use a switch for now.
func (r Robot) TranslateOrientation() (dx, dy int) {
    var nx, ny int

    switch r.orientation {
    case "N":
        nx, ny = 0, 1
    case "S":
        nx, ny = 0, -1
    case "E":
        nx, ny = 1, 0
    case "W":
        nx, ny = -1, 0
    }

    // We could have used `dx, dy` in the body of this
    // function instead of `nx, ny` with a bare `return`
    // but that can be confusing and IMHO explicit is better.
    return nx, ny
}

func (r *Robot) Forward() {
    dx, dy := r.TranslateOrientation()
    r.posX = r.posX + dx
    r.posY = r.posY + dy
}

func (r *Robot) TurnLeft() {
}

func (r *Robot) TurnRight() {
}

func (r Robot) OnScentedCell() bool {
    x, y := r.Position()
    return r.world.HasScent(x, y)
}
