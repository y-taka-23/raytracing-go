package raytracing

import "math"

type vector struct {
	x float64
	y float64
	z float64
}

func newVector(x, y, z float64) vector {
	return vector{x: x, y: y, z: z}
}

func (v vector) neg() vector {
	return newVector(-v.x, -v.y, -v.z)
}

func (v vector) add(w vector) vector {
	return newVector(v.x+w.x, v.y+w.y, v.z+w.z)
}

func (v vector) sub(w vector) vector {
	return newVector(v.x-w.x, v.y-w.y, v.z-w.z)
}

func (v vector) mul(t float64) vector {
	return newVector(t*v.x, t*v.y, t*v.z)
}

func (v vector) div(t float64) vector {
	return newVector(v.x/t, t*v.y/t, t*v.z/t)
}

func (v vector) dot(w vector) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v vector) cross(w vector) vector {
	return newVector(v.y*w.z-v.z*w.y, v.z*w.x-v.x*w.z, v.x*w.y-v.y*w.x)
}

func (v vector) norm() float64 {
	return v.dot(v)
}

func (v vector) length() float64 {
	return math.Sqrt(v.norm())
}

func (v vector) normalize() vector {
	return v.div(v.length())
}
