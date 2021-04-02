package raytracing

import (
	"math"
	"math/rand"
)

type Camera struct {
	lookFrom          Point
	horizontal        Vector
	vertical          Vector
	toLowerLeftCorner Vector
	u                 Vector
	v                 Vector
	lensRadius        float64
}

func NewCamera(lookFrom, lookAt Point, viewUp Vector,
	verticalFOV, aspectRatio, aperture, focusDistance float64) Camera {

	h := math.Tan(verticalFOV / 2)
	viewportHeight := 2 * h
	viewportWidth := viewportHeight * aspectRatio

	w := lookAt.to(lookFrom).normalize()
	u := viewUp.cross(w).normalize()
	v := w.cross(u)

	horizontal := u.mul(focusDistance * viewportWidth)
	vertical := v.mul(focusDistance * viewportHeight)
	toLowerLeftCorner := w.neg().mul(focusDistance).
		sub(horizontal.div(2)).
		sub(vertical.div(2))

	return Camera{
		lookFrom:          lookFrom,
		horizontal:        horizontal,
		vertical:          vertical,
		toLowerLeftCorner: toLowerLeftCorner,
		u:                 u,
		v:                 v,
		lensRadius:        aperture / 2.0,
	}
}

func (c Camera) castRay(s, t float64) ray {
	r := randomInUnitDisk().mul(c.lensRadius)
	offset := c.u.mul(r.x).add(c.v.mul(r.y))
	return newRay(
		origin().to(c.lookFrom).add(offset).point(),
		c.toLowerLeftCorner.
			add(c.horizontal.mul(s)).add(c.vertical.mul(t)).
			sub(offset),
	)
}

func randomInUnitDisk() Vector {
	x, y := 0.0, 0.0
	for true {
		x = 2*rand.Float64() - 1
		y = 2*rand.Float64() - 1
		if x*x+y*y < 1 {
			break
		}
	}
	return NewVector(x, y, 0.0)
}
