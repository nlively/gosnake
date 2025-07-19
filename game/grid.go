package game

type Grid struct {
	width  int
	height int
	points [][]bool
}

func NewGrid(width int, height int) *Grid {
	points := make([][]bool, width)
	for i := 0; i < width; i++ {
		points[i] = make([]bool, height)
	}

	return &Grid{width, height, points}
}

func (g *Grid) PlotPoints(points []Point) {
	for i := range points {
		g.PlotPoint(points[i])
	}
}

func (g *Grid) IsPointOutOfBounds(point Point) bool {
	if point.X < 0 || point.X >= g.width || point.Y < 0 || point.Y >= g.height {
		return true
	}
	return false
}

func (g *Grid) PlotPoint(point Point) {
	if g.IsPointOutOfBounds(point) {
		return
	}
	g.points[point.X][point.Y] = true
}

func (g *Grid) UnplotPoint(point Point) {
	if g.IsPointOutOfBounds(point) {
		return
	}
	g.points[point.X][point.Y] = false
}

func (g *Grid) IsPointFilled(point Point) bool {
	if g.IsPointOutOfBounds(point) {
		return false
	}
	return g.points[point.X][point.Y]
}
