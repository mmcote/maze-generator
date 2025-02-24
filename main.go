package main

import (
	"ca.michaelmauricejosephcote/maze/maze"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	m := maze.Generate(50, 30)
	m.Print()
}
