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
