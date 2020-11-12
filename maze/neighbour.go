package maze

type neighbour struct {
	wall bool
	cell *cell
}

func NewNeighbour(c *cell) *neighbour {
	return &neighbour{
		wall: true,
		cell: c,
	}
}
