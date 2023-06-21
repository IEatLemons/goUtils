package geometry

import (
	"math"
	"math/rand"
	"sort"
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

func CalculatePolygonArea(vertices Polygon) float64 {
	area := 0.0

	for i := 0; i < len(vertices.Vertices); i++ {
		j := (i + 1) % len(vertices.Vertices)
		area += vertices.Vertices[i].X * vertices.Vertices[j].Y
		area -= vertices.Vertices[j].X * vertices.Vertices[i].Y
	}

	area = math.Abs(area) / 2.0
	return area
}

// Define a function that implements Graham scanning and returns the convex hull of a set of points
func grahamScan(points []Point) []Point {
	// Gets the length of the point set
	n := len(points)
	// If the point set length is less than 3, return directly to the origin set
	if n < 3 {
		return points
	}
	// Find the point with the smallest Y coordinate, if there are more than one, take the one with the smallest X coordinate
	min := points[0]
	minIndeX := 0
	for i := 1; i < n; i++ {
		if points[i].Y < min.Y || (points[i].Y == min.Y && points[i].X < min.X) {
			min = points[i]
			minIndeX = i
		}
	}
	// Swap the smallest point to the first position
	points[0], points[minIndeX] = points[minIndeX], points[0]

	// Sort the remaining points by polar Angle from smallest to largest, and if the polar Angle is the same, sort by distance from nearest to far
	sort.Slice(points[1:], func(i, j int) bool {
		i++ // Since the slice starts with the second element, add one
		j++
		angle1 := polarAngle(min, points[i])
		angle2 := polarAngle(min, points[j])
		if angle1 == angle2 {
			return distance(min, points[i]) < distance(min, points[j])
		}
		return angle1 < angle2
	})

	// Initializes a stack to hold the points on the convex hull by stacking the first three points
	stack := make([]Point, n)
	stack[0] = points[0]
	stack[1] = points[1]
	stack[2] = points[2]

	// Initialize the top of the stack pointer to 2
	top := 2

	// Walk through the remaining points, judging each point
	for i := 3; i < n; i++ {
		// If the two points at the top of the stack and the current point form a right turn or collinear, the current point is not on the convex hull, and the top point is removed from the stack
		for top >= 1 && crossProduct(Point{stack[top-1].X - stack[top].X, stack[top-1].Y - stack[top].Y}, Point{points[i].X - stack[top].X, points[i].Y - stack[top].Y}) <= 0 {
			top--
		}
		// Otherwise, push the current point onto the stack
		top++
		stack[top] = points[i]
	}

	// Returns the point in the stack, that is, the point on the convex hull
	return stack[:top+1]
}

// Define a function that randomly generates some points
func GeneratePoints(n int, min, maX float64) []Point {
	// Initializes an empty slice to hold the generated points
	points := make([]Point, 0, n)
	// Initializes a map to remove the weight
	seen := make(map[Point]bool)
	// Loop n times, generating one point each time
	for i := 0; i < n; i++ {
		// X and Y coordinates are randomly generated in the range [min, maX]
		X := min + rand.Float64()*(maX-min)
		Y := min + rand.Float64()*(maX-min)
		// Create a point
		p := Point{X, Y}
		// If this point already exists, skip the loop
		if seen[p] {
			i--
			continue
		}
		// Otherwise, add the point to the slice and map
		points = append(points, p)
		seen[p] = true
	}
	// Returns the generated point
	return points
}

// Define a function that generates an irregular polygon and returns the vertex coordinates of the polygon
func GeneratePolYgon(n int, min, maX float64) []Point {
	// It randomly generates n points
	points := GeneratePoints(n, min, maX)
	// Use Graham scanning to find the convex hull of these points, i.e. the irregular polygon
	polYgon := grahamScan(points)
	// Returns the vertex coordinates of the polygon
	return polYgon
}
