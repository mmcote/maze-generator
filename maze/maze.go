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
		cells:  initializeMaze(h, w),
	}
}

// TODO: Candidate for removal
// - Need to remove reliance on this function when drawing the maze
// - Easiest way, keep track of root node (0, 0) in maze object
func (m *maze) getCell(x int, y int) *cell {
	return m.cells[x][y]
}

func _initializeWalls(above []*cell, below []*cell) {
	w := len(above)
	for i := 0; i < w - 1; i++ {
		above[i].walls[right] = NewWall(above[i + 1])
		above[i + 1].walls[left] = NewWall(above[i])
	}

	if below != nil {
		for i := 0; i < w; i++ {
			above[i].walls[down] = NewWall(below[i])
			below[i].walls[up] = NewWall(above[i])
		}
	}
}
func initializeMaze(h int, w int) [][]*cell {
	s := make([][]*cell, h)
	for i := range s {
		s[i] = make([]*cell, w)
		for j := range s[i] {
			s[i][j] = NewCell()
		}
	}

	for i := 0; i < h - 1; i++ {
		_initializeWalls(s[i], s[i + 1])
	}
	_initializeWalls(s[len(s) - 1], nil)

	return s
}

// figure out how to make `make` create a maze

func Generate(height int, width int) *maze {
	m := NewMaze(height, width)

	// stack for nodes to explore
	stack := make([]*cell, 0)

	// generate the coordinates of the initial cell
	_x := rand.Intn(width)
	_y := rand.Intn(height)

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

