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
    lost, ignored := false, false

    dx, dy := r.TranslateOrientation()
    // BUG: we update the position *before* checking if
    // we're lost or should ignore this cell.
    r.posX = r.posX + dx
    r.posY = r.posY + dy

    if r.posX < 0 || r.posY < 0 || r.posX >= r.world.BoundX || r.posY >= r.world.BoundY {
        // All robots get lost for now
        lost = true
    }

    return lost, ignored
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

func (r *Robot) DoCommand(c string) (bool, bool) {
    var lost, ignored bool

    switch c {
    case "F":
        lost, ignored = r.Forward()
    case "L":
        r.TurnLeft()
    case "R":
        r.TurnRight()
    // Ignore unknown commands or consider them a failure?
    default:
        lost, ignored = true, true
    }

    return lost, ignored
}

func (r Robot) Commands(c string) bool {
func (r *Robot) Commands(c string) bool {
    var lost bool

    for i:=0; i<len(c); i++ {
        // We don't care about ignored commands yet.
        lost, _ = r.DoCommand(c[i:i+1])
        // If we get lost, abort the command stream.
        // We could just `return true` here but maybe
        // there's cleanup, etc., we want to do after.
        if lost { break }
    }

    // Single return at the end can be cleaner code.
    return lost
}

func (r Robot) Report() string {
    return fmt.Sprintf("%d %d %s", r.posX, r.posY, r.orientation)
}
