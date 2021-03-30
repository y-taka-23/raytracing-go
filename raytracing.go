package raytracing

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"
)

const (
	aspectRatio = 16.0 / 9.0

	imageWidth  = 384
	imageHeight = int(imageWidth / aspectRatio)

	samplesPerPixel = 100
	maxDepth        = 50
)

func Run(stdout, stderr io.Writer) error {

	world := newHitters().
		add(newSphere(
			newPoint(0, 0, -1), 0.5,
			newLambertian(newColor(0.1, 0.2, 0.5)))).
		add(newSphere(
			newPoint(0, -100.5, -1), 100,
			newLambertian(newColor(0.8, 0.8, 0.0)))).
		add(newSphere(
			newPoint(1, 0, -1), 0.5,
			newMetal(newColor(0.8, 0.6, 0.2), 0.0))).
		add(newSphere(
			newPoint(-1, 0, -1), 0.5,
			newDielectric(1.5))).
		add(newSphere(
			newPoint(-1, 0, -1), -0.45,
			newDielectric(1.5)))

	cam := newCamera(newPoint(-2, 2, 1), newPoint(0, 0, -1), newVector(0, 1, 0),
		math.Pi/9.0, float64(imageWidth)/float64(imageHeight))

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	rand.Seed(time.Now().UnixNano())

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
