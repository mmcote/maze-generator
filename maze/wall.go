package maze

type Side int

const (
	up Side = iota
	down
	right
	left
)

type wall struct {
	side Side
	opposing *cell
}

func NewWall(s Side, c *cell) *wall {
	return &wall{
		side: s,
		opposing: c,
	}
}

func (w *wall) getOpposingCell() *cell {
	return w.opposing
}