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

    r.TurnRight()

    if r.Direction() != "E" {
        t.Fatalf("N+R should be E")
    }

    r.TurnRight()
    r.TurnRight()
    r.TurnRight()

    if r.Direction() != "N" {
        t.Fatalf("E+RRR should be N")
    }

    r.TurnRight()
    r.TurnLeft()

    if r.Direction() != "N" {
        t.Fatalf("N+LR should be N")
    }

    // Facing S now
    r.TurnRight()
    r.TurnRight()

    lost, ignored = r.Forward()
    if lost || ignored {
        t.Fatalf("Robot shouldn't get lost or ignore yet")
    }

    px, py = r.Position()
    if px != 1 || py != 1 {
        t.Fatalf("1,1 + FRRF should be 1,1")
    }

    lost, ignored = r.Forward()
    if lost || ignored {
        t.Fatalf("Robot shouldn't get lost or ignore yet")
    }

    px, py = r.Position()
    if px != 1 || py != 0 {
        t.Fatalf("1,1 S + F should be 1,0")
    }

    // We're about to step off the map.
    // Since we're the only robot, this should report 'LOST'
    lost, ignored = r.Forward()
    if !lost || ignored {
        t.Fatalf("Robot should be lost and command not ignored")
    }
}

// Clunky version of the first sample robot.
// We'll clean this up in a minute.
func TestSampleData(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(1, 1, "E", w)
    r.TurnRight()
    r.Forward()
    r.TurnRight()
    r.Forward()
    r.TurnRight()
    r.Forward()
    r.TurnRight()
    r.Forward()


    // Output should be "1 1 E"
    end := r.Report()
    if end != "1 1 E" {
        t.Fatalf("First robot failed to provide '1 1 E'")
    }
}

func TestCleanRobotOne(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(1, 1, "E", w)
    lost := r.Commands("RFRFRFRF")

    if lost {
        t.Fatalf("First robot does not get lost")
    }

    // Output should be "1 1 E"
    end := r.Report()
    if end != "1 1 E" {
        t.Fatalf("First robot failed to provide '1 1 E'")
    }
}

// Since Robot One doesn't get lost, we can check Robot Two
// without having to run Robot One first.
func TestCleanRobotTwo(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(3, 2, "N", w)
    lost := r.Commands("FRRFLLFFRRFLL")

    if !lost {
        t.Fatalf("Second robot does get lost")
    }

    // Output should be "3 3 N LOST"
    end := r.Report()
    if end != "3 3 N LOST" {
        t.Fatalf("Second robot failed to provide '3 3 N LOST': %s", end)

    }
}

// Test Robots 2+3
func TestCleanRobotTwoThree(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(3, 2, "N", w)
    lost := r.Commands("FRRFLLFFRRFLL")

    if !lost {
        t.Fatalf("Second robot does get lost")
    }

    // Output should be "3 3 N LOST"
    end := r.Report()
    if end != "3 3 N LOST" {
        t.Fatalf("Second robot failed to provide '3 3 N'")
    }

    r = NewRobot(0, 3, "W", w)
    lost = r.Commands("LLFFFLFLFL")

    if lost {
        t.Fatalf("Robot 3 does not get lost")
    }

    // Expecting '2 3 S'
    end = r.Report()
    if end != "2 3 S" {
        t.Fatalf("Robot 3 failed to report '2 3 S'")
    }
}

func TestStrings(t *testing.T) {
    w := NewWorldFromString("5 3")
    r := NewRobotFromString("1 1 E", w)
    lost := r.Commands("RFRFRFRF")

    if lost {
        t.Fatalf("First robot does not get lost")
    }

    // Output should be "1 1 E"
    end := r.Report()
    if end != "1 1 E" {
        t.Fatalf("First robot failed to provide '1 1 E'")
    }
}

