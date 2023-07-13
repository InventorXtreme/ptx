package main

import "math"

type V3 struct {
	X float64
	Y float64
	Z float64
}

func (v V3) Magnitude() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v V3) Normalize() V3 {
	nX := v.X / v.Magnitude()
	nY := v.Y / v.Magnitude()
	nZ := v.Z / v.Magnitude()
	return V3{nX, nY, nZ}
}

func V3Add(vector1 V3, vector2 V3) V3 {
	newx := vector1.X + vector2.X
	newy := vector1.Y + vector2.Y
	newz := vector1.Z + vector2.Z
	return V3{X: newx, Y: newy, Z: newz}
}

func V3Sub(vector1 V3, vector2 V3) V3 {
	newx := vector1.X - vector2.X
	newy := vector1.Y - vector2.Y
	newz := vector1.Z - vector2.Z
	return V3{X: newx, Y: newy, Z: newz}
}

func V3DotProduct(vector1 V3, vector2 V3) float64 {
	xprod := vector1.X * vector2.X
	yprod := vector1.Y * vector2.Y
	zprod := vector1.Z * vector2.Z

	var sum float64 = xprod + yprod + zprod
	return sum
}

func V3MultByFloat(vector V3, f float64) V3 {
	return V3{vector.X * f, vector.Y * f, vector.Z * f}
}
