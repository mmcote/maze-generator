package maze

import (
	"fmt"
	"math/rand"
	"strings"
)

type maze struct {
	height, width int
	cells [][]*cell
}

func NewMaze(h int, w int) *maze {
	return &maze{
		height: h,
		width:  w,
		cells:  initializeMazeCells(h, w),
	}
}

// TODO: Candidate for removal
// - Need to remove reliance on this function when drawing the maze
// - Easiest way, keep track of root node (0, 0) in maze object
func (m *maze) getCell(x int, y int) *cell {
	return m.cells[x][y]
}

func initializeMazeCells(h int, w int) [][]*cell {
	s := make([][]*cell, h)
	for i := range s {
		s[i] = make([]*cell, w)
		for j := range s[i] {
			s[i][j] = NewCell()
		}
	}

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			walls := &s[row][col].walls
			// Add neighbour above
			if row > 0 {
				*walls = append(*walls, NewWall(up, s[row-1][col]))
			}

			// Add neighbour below
			if row < (h - 1) {
				*walls = append(*walls, NewWall(down, s[row+1][col]))
			}


			// Add neighbour to the left
			if col > 0 {
				*walls = append(*walls, NewWall(left, s[row][col-1]))
			}

			// Add neighbour to the right
			if col < (w - 1) {
				*walls = append(*walls, NewWall(right, s[row][col+1]))
			}
		}
	}

	return s
}

// figure out how to make `make` create a maze

func Generate(height int, width int) *maze {
	m := NewMaze(height, width)

	// stack for nodes to explore
	stack := make([]*cell, 0)

	// generate the coordinates of the initial cell
	_x := int(rand.Intn(int(width)))
	_y := int(rand.Intn(int(height)))

	// TODO: Remove this, generate random
	_x = 0
	_y = 0

	// add the initial cell to the stack and mark it as visited
	cell := m.cells[_y][_x]
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
	for i := 0; i < m.width*2 + 1; i++ {
		fmt.Printf("|%d|", i % 10)
	}
	fmt.Println()
}

const wallSym = "|||"
const cellSym = "   "

// TODO: Clean up bro
func (m* maze) PrintMaze() {
	m._printRowBoarder()
	defer m._printRowBoarder()

	for row := 0; row < m.height; row++ {
		var cellLineBuilder strings.Builder
		cellLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row + 1) % 10))

		var wallLineBuilder strings.Builder
		wallLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row + 2) % 10))

		// Print in top-left to bottom-right fashion
		for col := 0; col < m.width; col++ {
			cell := m.getCell(row, col)
			cellLineBuilder.WriteString(cellSym)

			if col != m.width - 1 {
				if cell.hasWall(right) {
					cellLineBuilder.WriteString(wallSym)
				} else {
					cellLineBuilder.WriteString(cellSym)
				}
			}

			if row != m.height - 1 {
				if cell.hasWall(down) {
					wallLineBuilder.WriteString(wallSym)
				} else {
					wallLineBuilder.WriteString(cellSym)
				}
			}

			if col != m.width - 1 {
				wallLineBuilder.WriteString(wallSym)
			}
		}
		cellLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row + 1) % 10))
		wallLineBuilder.WriteString(fmt.Sprintf("|%d|", (2*row + 2) % 10))

		fmt.Println(cellLineBuilder.String())
		if row != m.height - 1 {
			fmt.Println(wallLineBuilder.String())
		}
	}
}

