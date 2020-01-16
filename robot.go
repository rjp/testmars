package main
import ("fmt")

type Robot struct {
    posX, posY int
    orientation string
    world World
    lost bool
}

// A Robot must have awareness of the World he's in.
// Physical analogy would be, e.g., a paper map.
func NewRobot(x int, y int, o string, w World) Robot {
    temp := Robot{posX: x, posY: y, orientation: o, world: w, lost: false}
    return temp
}

func NewRobotFromString(s string, w World) Robot {
    var x, y int
    var o string
    n, err := fmt.Sscanf(s, "%d %d %s", &x, &y, &o)
    if err != nil || n != 3 {
        panic("Parsing robot definition string")
    }
    return NewRobot(x, y, o, w)
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
    ignored := false

    // Work out our potential new position. Since we want
    // to drop scent if we get lost, we can't update our
    // position before we check the new one.
    dx, dy := r.TranslateOrientation()
    newX := r.posX + dx
    newY := r.posY + dy

    // FIXED: BoundX, BoundY are the coords of the corner, not the count of cells
    if newX < 0 || newY < 0 || newX > r.world.BoundX || newY > r.world.BoundY {
        // Ignore movements that might cause us to be `LOST` iff
        // we're standing on a scented cell.
        if r.OnScentedCell() {
            ignored = true
        } else {
            r.AddScent()
            r.lost = true
        }
    }

    // If we didn't get lost and we didn't ignore this command,
    // we can update our position.
    if !r.lost && !ignored {
        r.posX = newX
        r.posY = newY
    }

    return r.lost, ignored
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

// Drop a spray of scent on a cell to mark disaster.
func (r Robot) AddScent() {
    x, y := r.Position()
    r.world.AddScent(x, y)
}

// Check whether a cell is scented to avert disaster.
func (r Robot) OnScentedCell() bool {
    x, y := r.Position()
    return r.world.HasScent(x, y)
}

// Process a single command and return the outcome.
// We don't actually need to return `ignored` here
// because we never use it.
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
    // For now, let's consider them a lost robot (brain freeze.)
    default:
        lost, ignored = true, true
    }

    return lost, ignored
}

func (r *Robot) Commands(c string) bool {
    var lost bool

    for i:=0; i<len(c); i++ {
        // We don't care about ignored commands
        // because they don't interrupt the stream.
        lost, _ = r.DoCommand(c[i:i+1])

        // If we get lost, abort the command stream.
        // We could just `return true` here but maybe
        // there's cleanup, etc., we want to do after.
        if lost {
            r.lost = true
            break
        }
    }

    // Single return at the end can be cleaner code.
    return lost
}

func (r Robot) Report() string {
    var isLost string
    // If the robot was lost, we need a trailing " LOST" on the output.
    if r.lost {
        isLost = " LOST"
    }

    return fmt.Sprintf("%d %d %s%s", r.posX, r.posY, r.orientation, isLost)
}
