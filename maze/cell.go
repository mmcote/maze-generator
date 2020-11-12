package maze

import "math/rand"

type cell struct {
	visited    bool
	neighbours map[Side]*neighbour
}

func NewCell() *cell {
	return &cell{
		visited:    false,
		neighbours: make(map[Side]*neighbour, 0),
	}
}

func (c *cell) getRandomNeighbour() *cell {
	unvisited := make([]*neighbour, 0)
	for _, w := range c.neighbours {
		if !w.cell.visited {
			unvisited = append(unvisited, w)
		}
	}
	unvisitedCount := len(unvisited)
	if unvisitedCount == 0 {
		return nil
	}

	return unvisited[rand.Intn(unvisitedCount)].cell
}

func (c *cell) removeWall(n *cell) {
	for _, _n := range c.neighbours {
		if n == _n.cell {
			_n.wall = false
			return
		}
	}
}

func (c *cell) hasWall(s Side) bool {
	if n, ok := c.neighbours[s]; ok && n.wall {
		return true
	}

	return false
}
