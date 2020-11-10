package maze

import (
	"fmt"
	"math/rand"
	"strings"
)

type maze struct {
	root *cell
}

func NewMaze(h int, w int) *maze {
	return &maze{initializeMaze(h, w)}
}

func _initializeNeighbours(above []*cell, below []*cell) {
	w := len(above)
	for i := 0; i < w - 1; i++ {
		above[i].neighbours[right] = NewNeighbour(above[i + 1])
		above[i + 1].neighbours[left] = NewNeighbour(above[i])
	}

	if below != nil {
		for i := 0; i < w; i++ {
			above[i].neighbours[down] = NewNeighbour(below[i])
			below[i].neighbours[up] = NewNeighbour(above[i])
		}
	}
}
func initializeMaze(h int, w int) *cell {
	s := make([][]*cell, h)
	for i := range s {
		s[i] = make([]*cell, w)
		for j := range s[i] {
			s[i][j] = NewCell()
		}
	}

	for i := 0; i < h - 1; i++ {
		_initializeNeighbours(s[i], s[i + 1])
	}
	_initializeNeighbours(s[len(s) - 1], nil)

	return s[0][0]
}

func (m *maze) getRandomStartingCell(h int, w int) *cell {
	// generate the coordinates of the initial cell
	_x := rand.Intn(w)
	_y := rand.Intn(h)

	c := m.root
	for i := 0; i < _x; i++ {
		c = c.neighbours[right].cell
	}

	for i := 0; i < _y; i++ {
		c = c.neighbours[down].cell
	}

	return c
}

// figure out how to make `make` create a maze

func Generate(height int, width int) *maze {
	m := NewMaze(height, width)

	// stack for nodes to explore
	stack := make([]*cell, 0)

	// add the initial cell to the stack and mark it as visited
	cell := m.getRandomStartingCell(height, width)
	cell.visited = true

	stack = append(stack, cell)
	for len(stack) > 0 {
		l := len(stack)
		cell, stack = stack[l - 1], stack[:l - 1]

		// choose a random neighbour to explore
		neighbour := cell.getRandomNeighbour()
		if neighbour == nil {
			continue
		}

		stack = append(stack, cell)

		// clear the wall between the random neighbour and the original cell
		cell.removeWall(neighbour)
		neighbour.removeWall(cell)

		// mark the cell as visited and push it to the stack
		neighbour.visited = true
		stack = append(stack, neighbour)
	}

	return m
}

func (m* maze) _printRowBoarder() {
	col := 0
	c := m.root
	for c != nil {
		fmt.Printf("|%d||%d|", 2*col % 10, (2*col + 1) % 10)

		if _, ok := c.neighbours[right]; !ok {
			fmt.Printf("|%d|", (2*col + 2) % 10)
			break
		}

		c = c.neighbours[right].cell
		col++
	}
	fmt.Println()
}

const wallSym = "|||"
const cellSym = "   "

func _printRow(row int, c *cell) {
	// create cell line and draw left border
	var cellLineBuilder strings.Builder
	cellLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row+1)%10))

	// create wall line below draw wall line below
	var wallLineBuilder strings.Builder
	wallLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row+2)%10))

	for {
		cellLineBuilder.WriteString(cellSym)
		if _, ok := c.neighbours[right]; ok {
			if c.hasWall(right) {
				cellLineBuilder.WriteString(wallSym)
			} else {
				cellLineBuilder.WriteString(cellSym)
			}
		}

		if _, ok := c.neighbours[down]; ok {
			if c.hasWall(down) {
				wallLineBuilder.WriteString(wallSym)
			} else {
				wallLineBuilder.WriteString(cellSym)
			}
		}
		if _, ok := c.neighbours[right]; ok {
			wallLineBuilder.WriteString(wallSym)
		}

		if _, ok := c.neighbours[right]; !ok {
			break
		}

		c = c.neighbours[right].cell
	}

	cellLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row+1)%10))
	wallLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row+2)%10))

	fmt.Println(cellLineBuilder.String())
	if _, ok := c.neighbours[down]; ok {
		fmt.Println(wallLineBuilder.String())
	}
}

func (m* maze) PrintMaze() {
	m._printRowBoarder()
	defer m._printRowBoarder()

	row := 0
	c := m.root
	for {
		_printRow(row, c)

		if _, ok := c.neighbours[down]; !ok {
			break
		}

		c = c.neighbours[down].cell
		row++
	}
}

