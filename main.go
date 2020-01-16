package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// We'll be reusing these and `:=` is too scope'd
	var w World
	var r Robot

	// We're expecting some data on `STDIN`
	scanner := bufio.NewScanner(os.Stdin)

	// `world` -> `robot` -> `commands` -> `robot` ...
	parseState := "world"

	for scanner.Scan() {
		text := scanner.Text()

		// Ignore blank lines
		if text == "" {
			continue
		}

		switch parseState {
		case "world":
			// This will panic if it can't be done. There's a case
			// to be made that it should return an `err` but there's
			// nothing we can do to recover from an incorrect world
			// definition.
			w = NewWorldFromString(text)
			parseState = "robot"
		case "robot":
			// We can't get to `robot` without `world` which means
			// that we definitely have a defined `World` in `w`.
			// This also `panic`s if we can't create one but there
			// is a stronger case for returning `err` since, in theory,
			// we can recover and keep parsing the file until we come
			// across a working robot definition.
			r = NewRobotFromString(text, w)
			parseState = "commands"
		case "commands":
			// Same here - we must already have a defined `Robot`
			_ = r.Commands(text)
			fmt.Println(r.Report())
			parseState = "robot"
		}
	}

	// Not worth a `panic` since we might actually have succeeded in
	// having some robots explore before we get to `EOF`? Question
	// of personal taste, I think. Is it an error to create a world
	// and robot but not give it any commands?
	if parseState != "robot" {
		fmt.Println("Error in parsing")
	}
}
