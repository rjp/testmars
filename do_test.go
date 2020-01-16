package main

import (
	"testing"
)

func TestNewRobot(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(1, 1, "N", w)
    px, py := r.Position()

    // Haven't moved, should be where we started
    if px != 1 || py != 1 {
        t.Fatalf("Not where we put him")
    }

    // None of the cells are yet scented
    if r.OnScentedCell() {
        t.Fatalf("All cells are unscented")
    }

    lost, ignored := r.Forward()

    if lost || ignored {
        t.Fatalf("Robot shouldn't get lost or ignore this")
    }

    px, py = r.Position()

    // One step North should put us at 1,2
    if px != 1 || py != 2 {
        t.Fatalf("Move North")
    }

    r.TurnLeft()

    if r.Direction() != "W" {
        t.Fatalf("N+L should be W")
    }

    r.TurnLeft()
    r.TurnLeft()
    r.TurnLeft()

    if r.Direction() != "N" {
        t.Fatalf("W+LLL should be N")
    }

}

