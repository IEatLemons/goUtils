package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

// Define a function to compute the cross product of two points
func crossProduct(p1, p2 Point) float64 {
	return p1.X*p2.Y - p1.Y*p2.X
}

// Define a function that calculates the distance between two points
func distance(p1, p2 Point) float64 {
	dX := p1.X - p2.X
	dY := p1.Y - p2.Y
	return math.Sqrt(dX*dX + dY*dY)
}

// Define a function that calculates the polar Angle between two points
func polarAngle(p1, p2 Point) float64 {
	dX := p2.X - p1.X
	dY := p2.Y - p1.Y
	return math.Atan2(dY, dX)
}

// Define a function to determine whether two points are equal
// func equal(p1, p2 Point) bool {
// 	return p1.X == p2.X && p1.Y == p2.Y
// }
