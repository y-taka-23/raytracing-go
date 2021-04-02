package raytracing

type ray struct {
	origin    Point
	direction Vector
}

func newRay(p Point, v Vector) ray {
	return ray{origin: p, direction: v}
}

func (r ray) at(t float64) Point {
	return origin().to(r.origin).add(r.direction.mul(t)).point()
}
