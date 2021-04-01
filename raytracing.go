package raytracing

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"
)

const (
	aspectRatio = 4.0 / 3.0

	imageWidth  = 384 * 2
	imageHeight = int(imageWidth / aspectRatio)

	samplesPerPixel = 100
	maxDepth        = 50
)

func Run(stdout, stderr io.Writer) error {

	rand.Seed(time.Now().UnixNano())

	world := randomSpheres()

	lookFrom := newPoint(13, 2, 3)
	lookAt := newPoint(0, 0, 0)
	viewUp := newVector(0, 1, 0)
	aperture := 0.1
	distToFocus := 10.0
	cam := newCamera(lookFrom, lookAt, viewUp,
		math.Pi/9.0, aspectRatio, aperture, distToFocus)

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			fmt.Fprintf(stderr, "\rScanlines remaining: %d", j)
			color := newColor(0, 0, 0)
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				c := rayColor(cam.castRay(u, v), world, maxDepth)
				color = newColor(color.x+c.x, color.y+c.y, color.z+c.z)
			}
			writeColor(stdout, color, samplesPerPixel)
		}
	}
	fmt.Fprintln(stderr, "\nDone.")

	return nil
}

func rayColor(r ray, world hitters, depth int) color {

	if depth <= 0 {
		return newColor(0, 0, 0)
	}

	hr, ok := world.hit(r, 0.001, math.MaxFloat64)
	if !ok {
		unit := r.direction.normalize()
		x := 0.5 * (unit.y + 1)
		return newColor((1-x)*1.0+x*0.5, (1-x)*1.0+x*0.7, (1-x)*1.0+x*1.0)
	}

	scattered, ok := hr.material.scatter(hr)
	if !ok {
		return newColor(0, 0, 0)
	}
	att := hr.material.attenuation()
	c := rayColor(scattered, world, depth-1)
	return newColor(att.x*c.x, att.y*c.y, att.z*c.z)
}

func randomSpheres() hitters {

	world := newHitters().
		add(newSphere(
			newPoint(0, -1000, 0), 1000,
			newLambertian(newColor(0.5, 0.5, 0.5))))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {

			chooseMaterial := rand.Float64()
			center := newPoint(
				float64(a)+0.9*rand.Float64(),
				0.2,
				float64(b)+0.9*rand.Float64(),
			)

			if newPoint(0, 0.2, 0).to(center).length() > 0.9 &&
				newPoint(4, 0.2, 0).to(center).length() > 0.9 &&
				newPoint(-4, 0.2, 0).to(center).length() > 0.9 {
				if chooseMaterial < 0.8 {
					r1, g1, b1 := rand.Float64(), rand.Float64(), rand.Float64()
					r2, g2, b2 := rand.Float64(), rand.Float64(), rand.Float64()
					albedo := newColor(r1*r2, g1*g2, b1*b2)
					world = world.add(newSphere(center, 0.2, newLambertian(albedo)))
				} else if chooseMaterial < 0.95 {
					albedo := newColor(
						0.5*rand.Float64()+0.5,
						0.5*rand.Float64()+0.5,
						0.5*rand.Float64()+0.5,
					)
					fizz := rand.Float64() * 0.5
					world = world.add(newSphere(center, 0.2, newMetal(albedo, fizz)))
				} else {
					world = world.add(newSphere(center, 0.2, newDielectric(1.5)))
				}
			}
		}
	}

	world = world.
		add(newSphere(
			newPoint(0, 1, 0), 1,
			newDielectric(1.5))).
		add(newSphere(
			newPoint(-4, 1, 0), 1,
			newLambertian(newColor(0.4, 0.2, 0.1)))).
		add(newSphere(
			newPoint(4, 1, 0), 1,
			newMetal(newColor(0.7, 0.6, 0.5), 0)))

	return world
}
