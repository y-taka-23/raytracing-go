package raytracing

import (
	"math"
)

type hitter interface {
	hit(r ray, tMin, tMax float64) (hitRecord, bool)
}

type hitters struct {
	hitters []hitter
}

func newHitters() hitters {
	return hitters{hitters: []hitter{}}
}

func (hs hitters) add(h hitter) hitters {
	return hitters{hitters: append(hs.hitters, h)}
}

func (hs hitters) hit(r ray, tMin, tMax float64) (hitRecord, bool) {
	hitAnything := false
	closestT := math.MaxFloat64
	var closest hitRecord
	for _, h := range hs.hitters {
		if hr, ok := h.hit(r, tMin, tMax); ok {
			hitAnything = true
			if hr.t < closestT {
				closestT = hr.t
				closest = hr
			}
		}
	}
	return closest, hitAnything
}

type hitRecord struct {
	point  point
	normal vector
	t      float64
}

func newHitRecord(p point, v vector, t float64) hitRecord {
	return hitRecord{point: p, normal: v, t: t}
}

type sphere struct {
	center point
	radius float64
}

func newSphere(p point, r float64) sphere {
	return sphere{center: p, radius: r}
}

func (s sphere) hit(r ray, tMin, tMax float64) (hitRecord, bool) {

	v := s.center.to(r.origin)
	a := r.direction.norm()
	b := 2 * r.direction.dot(v)
	c := v.norm() - s.radius*s.radius

	disc := b*b - 4*a*c
	if disc < 0 {
		return hitRecord{}, false
	}

	tClose := (-b - math.Sqrt(disc)) / (2 * a)
	if tMin < tClose && tClose < tMax {
		p := r.at(tClose)
		n := s.center.to(p).div(s.radius)
		return newHitRecord(p, n, tClose), true
	}

	tFar := (-b + math.Sqrt(disc)) / (2 * a)
	if tMin < tFar && tFar < tMax {
		p := r.at(tFar)
		n := s.center.to(p).div(s.radius)
		return newHitRecord(p, n, tFar), true
	}

	return hitRecord{}, false
}
