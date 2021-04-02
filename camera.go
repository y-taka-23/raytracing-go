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

type CameraConfig struct {
	aspectRatio   float64
	lookFrom      Point
	lookAt        Point
	viewUp        Vector
	verticalFOV   float64
	aperture      float64
	focusDistance float64
}

type CameraConfigOption func(*CameraConfig)

func NewCameraConfig(aspectRatio float64, opts ...CameraConfigOption) CameraConfig {
	cfg := CameraConfig{
		aspectRatio:   aspectRatio,
		lookFrom:      origin(),
		lookAt:        NewPoint(0, 0, -1),
		viewUp:        NewVector(0, 1, 0),
		verticalFOV:   math.Pi / 2.0,
		aperture:      0.0,
		focusDistance: 1.0,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

func WithAspectRatio(ratio float64) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.aspectRatio = ratio
	}
}

func WithPointOfView(from, to Point) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.lookFrom = from
		cfg.lookAt = to
	}
}

func WithViewUp(vup Vector) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.viewUp = vup
	}
}

func WithVerticalFOV(rad float64) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.verticalFOV = rad
	}
}

func WithAperture(aperture float64) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.aperture = aperture
	}
}

func WithFocusDistance(dist float64) CameraConfigOption {
	return func(cfg *CameraConfig) {
		cfg.focusDistance = dist
	}
}

func NewCamera(cfg CameraConfig) Camera {

	h := math.Tan(cfg.verticalFOV / 2)
	viewportHeight := 2 * h
	viewportWidth := viewportHeight * cfg.aspectRatio

	w := cfg.lookAt.to(cfg.lookFrom).normalize()
	u := cfg.viewUp.cross(w).normalize()
	v := w.cross(u)

	horizontal := u.mul(cfg.focusDistance * viewportWidth)
	vertical := v.mul(cfg.focusDistance * viewportHeight)
	toLowerLeftCorner := w.neg().mul(cfg.focusDistance).
		sub(horizontal.div(2)).
		sub(vertical.div(2))

	c := Camera{
		lookFrom:          cfg.lookFrom,
		horizontal:        horizontal,
		vertical:          vertical,
		toLowerLeftCorner: toLowerLeftCorner,
		u:                 u,
		v:                 v,
		lensRadius:        cfg.aperture / 2.0,
	}

	return c
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
