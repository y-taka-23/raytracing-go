package raytracing

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{x: x, y: y, z: z}
}

func (v Vector) neg() Vector {
	return NewVector(-v.x, -v.y, -v.z)
}

func (v Vector) add(w Vector) Vector {
	return NewVector(v.x+w.x, v.y+w.y, v.z+w.z)
}

func (v Vector) sub(w Vector) Vector {
	return NewVector(v.x-w.x, v.y-w.y, v.z-w.z)
}

func (v Vector) mul(t float64) Vector {
	return NewVector(t*v.x, t*v.y, t*v.z)
}

func (v Vector) div(t float64) Vector {
	return NewVector(v.x/t, v.y/t, v.z/t)
}

func (v Vector) dot(w Vector) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v Vector) cross(w Vector) Vector {
	return NewVector(v.y*w.z-v.z*w.y, v.z*w.x-v.x*w.z, v.x*w.y-v.y*w.x)
}

func (v Vector) norm() float64 {
	return v.dot(v)
}

func (v Vector) length() float64 {
	return math.Sqrt(v.norm())
}

func (v Vector) normalize() Vector {
	return v.div(v.length())
}

type Point Vector

func NewPoint(x, y, z float64) Point {
	return Point(NewVector(x, y, z))
}

func origin() Point {
	return NewPoint(0, 0, 0)
}

func (p Point) to(q Point) Vector {
	return NewVector(q.x-p.x, q.y-p.y, q.z-p.z)
}

func (v Vector) point() Point {
	return Point(v)
}
