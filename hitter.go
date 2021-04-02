package raytracing

import (
	"math"
)

type Hitter interface {
	hit(r ray, tMin, tMax float64) (hitRecord, bool)
}

type Scene struct {
	hitters []Hitter
}

func NewScene() *Scene {
	return &Scene{hitters: []Hitter{}}
}

func (s *Scene) Add(h Hitter) *Scene {
	s.hitters = append(s.hitters, h)
	return s
}

func (s Scene) hit(r ray, tMin, tMax float64) (hitRecord, bool) {
	hitAnything := false
	closestT := math.MaxFloat64
	var closest hitRecord
	for _, h := range s.hitters {
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
	point    Point
	normal   Vector
	incident ray
	t        float64
	material Material
}

func newHitRecord(p Point, v Vector, r ray, t float64, m Material) hitRecord {
	return hitRecord{point: p, normal: v, incident: r, t: t, material: m}
}

type sphere struct {
	center   Point
	radius   float64
	material Material
}

func NewSphere(p Point, r float64, m Material) Hitter {
	return sphere{center: p, radius: r, material: m}
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
		return newHitRecord(p, n, r, tClose, s.material), true
	}

	tFar := (-b + math.Sqrt(disc)) / (2 * a)
	if tMin < tFar && tFar < tMax {
		p := r.at(tFar)
		n := s.center.to(p).div(s.radius)
		return newHitRecord(p, n, r, tFar, s.material), true
	}

	return hitRecord{}, false
}
