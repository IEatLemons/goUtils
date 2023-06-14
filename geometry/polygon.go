package geometry

import (
	"math/rand"
	"time"
)

type Polygon struct {
	Vertices []Point
}

func GenerateRandomPolygon(numVertices int, minX, maxX, minY, maxY float64) *Polygon {
	rand.Seed(time.Now().UnixNano())

	polygon := &Polygon{}
	for i := 0; i < numVertices; i++ {
		x := rand.Float64()*(maxX-minX) + minX
		y := rand.Float64()*(maxY-minY) + minY
		polygon.Vertices = append(polygon.Vertices, Point{X: x, Y: y})
	}

	return polygon
}

// Determine if the point is inside the polygon
func isPointInPolygon(point Point, polygon []Point) bool {
	// Use the ray method to determine if the point is inside the polygon
	intersectCount := 0
	for i := 0; i < len(polygon); i++ {
		current := polygon[i]
		next := polygon[(i+1)%len(polygon)]

		// Determine whether the point intersects the sides of the polygon
		if (current.Y > point.Y) != (next.Y > point.Y) &&
			point.X < (next.X-current.X)*(point.Y-current.Y)/(next.Y-current.Y)+current.X {
			intersectCount++
		}
	}

	// Determine the parity of the number of intersection points
	return intersectCount%2 == 1
}

func Inclusion(polygon1, polygon2 []Point) bool {
	isContained := true
	for _, point := range polygon2 {
		if !isPointInPolygon(point, polygon1) {
			isContained = false
			break
		}
	}
	return isContained
}
