package raytracing

import (
	"math"
)

type ray struct {
	origin    point
	direction vector
}

func newRay(p point, v vector) ray {
	return ray{origin: p, direction: v}
}

func (r ray) at(t float64) point {
	return r.origin.vector().add(r.direction.mul(t)).point()
}

func (r ray) hitSphere(center point, radius float64) (float64, bool) {
	v := center.to(r.origin)
	a := r.direction.norm()
	b := 2 * r.direction.dot(v)
	c := v.norm() - radius*radius
	disc := b*b - 4*a*c
	if disc < 0 {
		return 0, false
	} else {
		return (-b - math.Sqrt(disc)) / (2 * a), true
	}
}
