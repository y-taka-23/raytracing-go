package raytracing

type ray struct {
	origin    point
	direction vector
}

func newRay(p point, v vector) ray {
	return ray{origin: p, direction: v}
}

func (r ray) at(t float64) vector {
	return r.origin.vector().add(r.direction.mul(t))
}

func (r ray) hitSphere(center point, radius float64) bool {
	v := center.to(r.origin)
	a := r.direction.norm()
	b := 2 * r.direction.dot(v)
	c := v.norm() - radius*radius
	return b*b-4*a*c > 0
}
