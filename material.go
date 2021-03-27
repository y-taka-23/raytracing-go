package raytracing

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(hr hitRecord) (ray, bool)
	attenuation() color
}

type lambertian struct {
	albedo color
}

func newLambertian(c color) material {
	return lambertian{albedo: c}
}

func (l lambertian) scatter(hr hitRecord) (ray, bool) {
	phi := 2 * math.Pi * rand.Float64()
	z := 2*rand.Float64() - 1
	r := (1 - z*z)
	random := newVector(r*math.Cos(phi), r*math.Sin(phi), z)
	return newRay(hr.point, hr.normal.add(random)), true
}

func (l lambertian) attenuation() color {
	return l.albedo
}

type metal struct {
	albedo color
	fuzz   float64
}

func newMetal(c color, f float64) material {
	if f <= 0 {
		f = 0
	}
	if f >= 1 {
		f = 1
	}
	return metal{albedo: c, fuzz: f}
}

func (m metal) scatter(hr hitRecord) (ray, bool) {
	v := hr.incident.direction
	w := hr.normal.mul(v.neg().dot(hr.normal))
	f := randomInUnitSphere().mul(m.fuzz)
	ref := v.add(w.mul(2)).add(f)
	if ref.dot(hr.normal) < 0 {
		return ray{}, false
	}
	return newRay(hr.point, ref), true
}

func randomInUnitSphere() vector {
	x, y, z := 0.0, 0.0, 0.0
	for true {
		x = 2*rand.Float64() - 1
		y = 2*rand.Float64() - 1
		z = 2*rand.Float64() - 1
		if x*x+y*y+z*z < 1 {
			break
		}
	}
	return newVector(x, y, z)
}

func (m metal) attenuation() color {
	return m.albedo
}
