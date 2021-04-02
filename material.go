package raytracing

import (
	"math"
	"math/rand"
)

type Material interface {
	scatter(hr hitRecord) (ray, bool)
	attenuation() Color
}

type lambertian struct {
	albedo Color
}

func NewLambertian(c Color) Material {
	return lambertian{albedo: c}
}

func (l lambertian) scatter(hr hitRecord) (ray, bool) {
	phi := 2 * math.Pi * rand.Float64()
	z := 2*rand.Float64() - 1
	r := (1 - z*z)
	random := NewVector(r*math.Cos(phi), r*math.Sin(phi), z)
	return newRay(hr.point, hr.normal.add(random)), true
}

func (l lambertian) attenuation() Color {
	return l.albedo
}

type metal struct {
	albedo    Color
	fuzziness float64
}

func NewMetal(c Color, f float64) Material {
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

func randomInUnitSphere() Vector {
	x, y, z := 0.0, 0.0, 0.0
	for true {
		x = 2*rand.Float64() - 1
		y = 2*rand.Float64() - 1
		z = 2*rand.Float64() - 1
		if x*x+y*y+z*z < 1 {
			break
		}
	}
	return NewVector(x, y, z)
}

func (m metal) attenuation() Color {
	return m.albedo
}

type dielectric struct {
	refractiveIndex float64
}

func NewDielectric(idx float64) Material {
	return dielectric{refractiveIndex: idx}
}

func (d dielectric) scatter(hr hitRecord) (ray, bool) {

	var relIdx float64
	if hr.incident.direction.dot(hr.normal) < 0 {
		relIdx = 1.0 / d.refractiveIndex
	} else {
		relIdx = d.refractiveIndex
	}

	in := hr.incident.direction
	orthogonal := in.sub(hr.normal.mul(in.dot(hr.normal))).mul(relIdx)

	if in.norm() < orthogonal.norm() {
		reflected := reflect(hr.incident.direction, hr.normal)
		return newRay(hr.point, reflected), true
	}

	if rand.Float64() < schlick(in, hr.normal, relIdx) {
		reflected := reflect(hr.incident.direction, hr.normal)
		return newRay(hr.point, reflected), true
	}

	var parallel Vector
	if hr.incident.direction.dot(hr.normal) < 0 {
		parallel = hr.normal.mul(-math.Sqrt(in.norm() - orthogonal.norm()))
	} else {
		parallel = hr.normal.mul(math.Sqrt(in.norm() - orthogonal.norm()))
	}

	refracted := orthogonal.add(parallel)
	return newRay(hr.point, refracted), true
}

func (d dielectric) attenuation() Color {
	return NewColor(1, 1, 1)
}

func reflect(in Vector, normal Vector) Vector {
	parallel := normal.mul(in.neg().dot(normal))
	return in.add(parallel.mul(2))
}

func schlick(in, normal Vector, relIdx float64) float64 {
	cos := in.dot(normal) / in.length()
	if cos < 0 {
		cos = -cos
	}
	r := (relIdx - 1) / (relIdx + 1)
	r0 := r * r
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
