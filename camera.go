package raytracing

import (
	"math"
	"math/rand"
)

type camera struct {
	lookFrom          point
	horizontal        vector
	vertical          vector
	toLowerLeftCorner vector
	u                 vector
	v                 vector
	lensRadius        float64
}

func newCamera(lookFrom, lookAt point, viewUp vector,
	verticalFOV, aspectRatio, aperture, focusDistance float64) camera {

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

	return camera{
		lookFrom:          lookFrom,
		horizontal:        horizontal,
		vertical:          vertical,
		toLowerLeftCorner: toLowerLeftCorner,
		u:                 u,
		v:                 v,
		lensRadius:        aperture / 2.0,
	}
}

func (c camera) castRay(s, t float64) ray {
	r := randomInUnitDisk().mul(c.lensRadius)
	offset := c.u.mul(r.x).add(c.v.mul(r.y))
	return newRay(
		origin().to(c.lookFrom).add(offset).point(),
		c.toLowerLeftCorner.
			add(c.horizontal.mul(s)).add(c.vertical.mul(t)).
			sub(offset),
	)
}

func randomInUnitDisk() vector {
	x, y := 0.0, 0.0
	for true {
		x = 2*rand.Float64() - 1
		y = 2*rand.Float64() - 1
		if x*x+y*y < 1 {
			break
		}
	}
	return newVector(x, y, 0.0)
}
