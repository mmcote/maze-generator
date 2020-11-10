package maze

type Side int

const (
	up Side = iota
	down
	right
	left
)

type wall struct {
	present bool
	cell    *cell
}

func NewWall(c *cell) *wall {
	return &wall{
		present: true,
		cell: c,
	}
}

type cell struct {
	visited bool
	walls map[Side]*wall
}

func NewCell() *cell {
	return &cell{
		visited: false,
		walls:   make(map[Side]*wall, 0),
	}
}

func (c *cell) getRandomNeighbour() *cell {
	for _, w := range c.walls {
		if !w.cell.visited {
			return w.cell
		}
	}

	return nil
}

func (c *cell) removeWall(n *cell) {
	for _, wall := range c.walls {
		if n == wall.cell {
			wall.present = false
			return
		}
	}
}

func (c* cell) hasWall(s Side) bool {
	if w, ok := c.walls[s]; ok && w.present{
		return true
	}

	return false
}
