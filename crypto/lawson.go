package crypto

import (
	"math"
)

// Point is a 2D point
type Point struct {
	X, Y float64
}

// Triangle is a triangle with three points
type Triangle struct {
	A, B, C Point
}

// Triangulation is a list of triangles
type Triangulation []Triangle

// SuperTriangle returns a triangle that contains all points in the given set
func SuperTriangle(points []Point) Triangle {
	// find the bounding box of the points
	minX := math.Inf(1)
	maxX := math.Inf(-1)
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	// create a triangle that is 10 times larger than the bounding box
	dx := maxX - minX
	dy := maxY - minY
	midX := (minX + maxX) / 2
	midY := (minY + maxY) / 2
	return Triangle{
		Point{midX - 10*dx, midY - dy},
		Point{midX, midY + 10*dy},
		Point{midX + 10*dx, midY - dy},
	}
}

// InCircle returns true if the point p is inside the circumcircle of the triangle t
func InCircle(p Point, t Triangle) bool {
	ax := t.A.X - p.X
	ay := t.A.Y - p.Y
	bx := t.B.X - p.X
	by := t.B.Y - p.Y
	cx := t.C.X - p.X
	cy := t.C.Y - p.Y

	return (ax*ax+ay*ay)*(bx*cy-by*cx)-(bx*bx+by*by)*(ax*cy-ay*cx)+(cx*cx+cy*cy)*(ax*by-ay*bx) > 0
}

// Lawson returns the Delaunay triangulation of the given points using Lawson's algorithm
func Lawson(points []Point) Triangulation {
	// initialize the triangulation with a super triangle that contains all points
	super := SuperTriangle(points)
	triangles := Triangulation{super}

	for _, p := range points {
		var edges []Edge // store the edges of the polygonal hole

		var newTriangles Triangulation // store the new triangles

		for i, t := range triangles {
			if InCircle(p, t) {
				// add the edges of this triangle to the edge list
				edges = append(edges, Edge{t.A, t.B}, Edge{t.B, t.C}, Edge{t.C, t.A})
				continue // skip this triangle
			}
			newTriangles = append(newTriangles, triangles[i]) // keep this triangle
		}

		for i := 0; i < len(edges); i++ {
			duplicate := false // flag to check if this edge is duplicated
			for j := i + 1; j < len(edges); j++ {
				if edges[i].Equals(edges[j]) {
					duplicate = true // this edge appears twice, so it is not part of the convex hull
					break            // no need to check further
				}
			}
			if !duplicate {
				// form a new triangle with this edge and the point p
				newTriangles = append(newTriangles, Triangle{edges[i].P1, edges[i].P2, p})
			}
		}

		triangles = newTriangles // update the triangulation

	}

	var result Triangulation // store the final result

	for _, t := range triangles {
		if !t.HasVertex(super.A) && !t.HasVertex(super.B) && !t.HasVertex(super.C) {
			result = append(result, t) // only keep triangles that are not part of the super triangle
		}
	}

	return result

}

// Edge is an edge with two points
type Edge struct {
	P1, P2 Point
}

// Equals returns true if the edge e is equal to the edge f (ignoring the order of points)
func (e Edge) Equals(f Edge) bool {
	return (e.P1 == f.P1 && e.P2 == f.P2) || (e.P1 == f.P2 && e.P2 == f.P1)
}

// HasVertex returns true if the point p is one of the vertices of the triangle t
func (t Triangle) HasVertex(p Point) bool {
	return t.A == p || t.B == p || t.C == p
}
