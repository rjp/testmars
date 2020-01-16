package main

import (
	"testing"
)

func TestNewRobot(t *testing.T) {
    w := NewWorld(5, 3)
    r := NewRobot(1, 1, "E", w)
    px, py := r.Position()
    if px != 1 || py != 1 {
        t.Fail()
    }
}

