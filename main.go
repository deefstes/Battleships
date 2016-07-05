package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Point represents an x-y coordinate
type Point struct {
	X, Y int
}

// NewPoint returns a new instance of Point
func NewPoint(x, y int) *Point {
	point := Point{x, y}
	return &point
}

// CellState defines the different states a cell can be in
type CellState int

const (
	unknown = iota
	miss    = iota
	hit     = iota
	sunk    = iota
)

// Cell is used to represent a single location on the grid
type Cell struct {
	location Point
	state    CellState
}

// GetCellDisplay returns a string representation for a cell
func (c Cell) GetCellDisplay() string {
	var retval string

	switch c.state {
	case unknown:
		retval = "."
	case miss:
		retval = "x"
	case hit:
		retval = "X"
	case sunk:
		retval = "*"
	}

	return retval
}

// NewCell return a new instance of Cell
func NewCell(location Point, state CellState) *Cell {
	cell := Cell{location, state}

	return &cell
}

// Grid contains the playing field
type Grid struct {
	cells [][]Cell
}

func create(x, y int) *Grid {
	grid := make([][]Cell, y)
	for cnty := range grid {
		grid[y] = make([]Cell, x)

		for cntx := range grid[cnty] {
			grid[cnty][cntx] = *NewCell(*NewPoint(x, y), unknown)
		}
	}

	return &Grid{grid}
}

// Game contains all the information pertaining to the current game
type Game struct {
	grid Grid
}

func (g Game) init(filename string) {
	lines, err := readFile("battleships.cfg")
	if err != nil {
		fmt.Println("Error reading battleships.cfg: %s", err)
	}

	for index, element := range lines {
		//fmt.Println(index, element)
		var x, y int
		if index == 0 {
			s := strings.Split(element, ",")
			x, _ = strconv.Atoi(s[0])
			y, _ = strconv.Atoi(s[1])
			fmt.Println(x, y)
			newgrid := create(x, y)
			g.grid = *newgrid
			//g.grid.cells = make([][]Cell, x*y)
		}
	}
}

// readFile reads a whole file into memory
// and returns a slice of its lines.
func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Entry point
func main() {
	fmt.Println("Battleships Solver")
	fmt.Println()

	var g Game
	g.init("battleships.bss")
}
