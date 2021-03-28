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
	albedo    color
	fuzziness float64
}

func newMetal(c color, f float64) material {
	if f <= 0 {
		f = 0
	}
	if f >= 1 {
		f = 1
	}
	return metal{albedo: c, fuzziness: f}
}

func (m metal) scatter(hr hitRecord) (ray, bool) {
	f := randomInUnitSphere().mul(m.fuzziness)
	reflected := reflect(hr.incident.direction, hr.normal).add(f)
	if reflected.dot(hr.normal) < 0 {
		return ray{}, false
	}
	return newRay(hr.point, reflected), true
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

func reflect(in vector, normal vector) vector {
	parallel := normal.mul(in.neg().dot(normal))
	return in.add(parallel.mul(2))
}
