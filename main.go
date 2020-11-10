package main

import (
	"ca.michaelmauricejosephcote/maze/maze"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	m := maze.Generate(5, 30)
	m.PrintMaze()
	fmt.Sprintln("%t", m)
}
