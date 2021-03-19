package raytracing

type camera struct {
	position       point
	aspectRatio    float64
	viewportHeight float64
	viewportWidth  float64
	focalLength    float64
}

func defaultCamera() camera {
	ratio := 16.0 / 9.0
	height := 2.0
	return camera{
		position:       origin(),
		aspectRatio:    ratio,
		viewportWidth:  height * ratio,
		viewportHeight: height,
		focalLength:    1,
	}
}

func (c camera) castRay(u, v float64) ray {

	horizontal := newVector(c.viewportWidth, 0, 0)
	vertical := newVector(0, c.viewportHeight, 0)
	toLowerLeftCorner := newVector(-c.viewportWidth/2, -c.viewportHeight/2, -c.focalLength)

	return newRay(
		c.position,
		toLowerLeftCorner.add(horizontal.mul(u)).add(vertical.mul(v)),
	)
}
