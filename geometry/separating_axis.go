package geometry

import "math"

func SeparatingAxis(polygon1, polygon2 Polygon) bool {
	axes := append(getAxes(polygon1), getAxes(polygon2)...)
	for _, axis := range axes {
		if !overlapOnAxis(polygon1, polygon2, axis) {
			return false
		}
	}
	return true
}

func getAxes(polygon Polygon) []Point {
	axes := make([]Point, len(polygon.Vertices))
	for i := 0; i < len(polygon.Vertices); i++ {
		current := polygon.Vertices[i]
		next := polygon.Vertices[(i+1)%len(polygon.Vertices)]
		axis := Point{Y: current.X - next.X, X: -(current.Y - next.Y)}
		axes[i] = normalize(axis)
	}
	return axes
}

func normalize(vector Point) Point {
	length := math.Sqrt(vector.X*vector.X + vector.Y*vector.Y)
	return Point{X: vector.X / length, Y: vector.Y / length}
}

func overlapOnAxis(polygon1, polygon2 Polygon, axis Point) bool {
	min1, max1 := project(polygon1, axis)
	min2, max2 := project(polygon2, axis)
	return max1 >= min2 && max2 >= min1
}

func project(polygon Polygon, axis Point) (min, max float64) {
	min = dotProduct(polygon.Vertices[0], axis)
	max = min
	for _, vertex := range polygon.Vertices[1:] {
		p := dotProduct(vertex, axis)
		if p < min {
			min = p
		} else if p > max {
			max = p
		}
	}
	return min, max
}

func dotProduct(point, axis Point) float64 {
	return point.X*axis.X + point.Y*axis.Y
}
