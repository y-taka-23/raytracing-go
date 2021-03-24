package raytracing

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(hr hitRecord) ray
	attenuation() color
}

type lambertian struct {
	albedo color
}

func newLambertian(c color) material {
	return lambertian{albedo: c}
}

func (l lambertian) scatter(hr hitRecord) ray {
	phi := 2 * math.Pi * rand.Float64()
	z := 2*rand.Float64() - 1
	r := (1 - z*z)
	random := newVector(r*math.Cos(phi), r*math.Sin(phi), z)
	return newRay(hr.point, hr.normal.add(random))
}

func (l lambertian) attenuation() color {
	return l.albedo
}

type metal struct {
	albedo color
}

func newMetal(c color) material {
	return metal{albedo: c}
}

func (m metal) scatter(hr hitRecord) ray {
	v := hr.incident.direction
	w := hr.normal.mul(v.neg().dot(hr.normal))
	return newRay(hr.point, v.add(w.mul(2)))
}

func (m metal) attenuation() color {
	return m.albedo
}
