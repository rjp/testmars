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

// Moving forward can have different results:
// 1. YES, moved forward ok
// 2. LOST, fell off the edge of the map
// 3. IGNORED, scented cell
// To which end we return two bools: lost, ignored
func (r *Robot) Forward() (bool, bool) {
    dx, dy := r.TranslateOrientation()
    r.posX = r.posX + dx
    r.posY = r.posY + dy

    return false, false
}

// This is clunky but we're writing simple code.
// A lookup table or matrix would be 'better'.
// Another approach would be to treat orientation
// as an `int mod 4`, then `TurnLeft` is simply `o-1`
// and `TurnRight` is `o+1`.
func (r *Robot) TurnLeft() {
    switch r.orientation {
    case "N":
        r.orientation = "W"
    case "E":
        r.orientation = "N"
    case "S":
        r.orientation = "E"
    case "W":
        r.orientation = "S"
    }
}

// Rather than do the `switch` again, we'll cheat
// and turn left three times.
func (r *Robot) TurnRight() {
    r.TurnLeft()
    r.TurnLeft()
    r.TurnLeft()
}

// `Direction` is shorter and easier to parse for
// people than `orientation`.
func (r Robot) Direction() string {
    return r.orientation
}

func (r Robot) OnScentedCell() bool {
    x, y := r.Position()
    return r.world.HasScent(x, y)
}
