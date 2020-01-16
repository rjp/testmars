package main

import (
    "fmt"
)

func main() {
    w := NewWorld(5, 3)
    fmt.Printf("%#v\n", w.HasScent(1,1))
    r := NewRobot(1, 1, "E", w)
    fmt.Printf("%#v\n", r)
}
