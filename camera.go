package raytracing

import (
	"math"
)

type camera struct {
	lookFrom          point
	horizontal        vector
	vertical          vector
	toLowerLeftCorner vector
}

func newCamera(lookFrom, lookAt point, viewUp vector, verticalFOV, aspectRatio float64) camera {

	h := math.Tan(verticalFOV / 2)
	viewportHeight := 2 * h
	viewportWidth := viewportHeight * aspectRatio

	w := lookAt.to(lookFrom).normalize()
	u := viewUp.cross(w).normalize()
	v := w.cross(u)

	horizontal := u.mul(viewportWidth)
	vertical := v.mul(viewportHeight)
	toLowerLeftCorner := w.neg().
		sub(horizontal.div(2)).
		sub(vertical.div(2))

	return camera{
		lookFrom:          lookFrom,
		horizontal:        horizontal,
		vertical:          vertical,
		toLowerLeftCorner: toLowerLeftCorner,
	}
}

func (c camera) castRay(u, v float64) ray {
	return newRay(
		c.lookFrom,
		c.toLowerLeftCorner.add(c.horizontal.mul(u)).add(c.vertical.mul(v)),
	)
}
