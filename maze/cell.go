package maze

import "math/rand"

// Walls are represented by edges between nodes
// Structure of the maze is represented by the 2D array
type cell struct {
	x, y int
	visited bool
	walls []*wall
}

func NewCell(x int, y int) cell {
	return cell{
		x: x,
		y: y,
		walls: make([]*wall, 0),
	}
}

func (c *cell) getRandomNeighbour() *cell {
	if len(c.walls) == 0 {
		return nil
	}

	index := rand.Intn(len(c.walls))

	// calculate the index that would result in all neighbours being visited
	// index = 0, len(c.walls) = 4, (0 + 4 - 1) % 4 = 3
	// index = 1, len(c.walls) = 4, (1 + 4 - 1) % 4 = 0
	// index = 2, len(c.walls) = 4, (2 + 4 - 1) % 4 = 1
	// index = 3, len(c.walls) = 4, (3 + 4 - 1) % 4 = 2
	finishIndex := (index + len(c.walls) - 1) % len(c.walls)
	for {
		neighbour := c.walls[index].getOpposingCell()
		if !neighbour.visited {
			return neighbour
		}

		if index == finishIndex {
			break
		}

		index = (index + 1) % len(c.walls)
	}

	return nil
}

func (c *cell) visit() {
	c.visited = true
}

func (c *cell) removeWall(n *cell) {
	for i, wall := range c.walls {
		if n == wall.getOpposingCell() {
			c.walls[i] = c.walls[len(c.walls) - 1]
			c.walls[len(c.walls) - 1] = nil
			c.walls = c.walls[:len(c.walls) - 1]

			return
		}
	}
}

func (c* cell) hasWall(side Side) bool {
	for _, wall := range c.walls {
		if side == wall.side {
			return true
		}
	}

	return false
}
