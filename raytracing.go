package raytracing

import (
	"fmt"
	"io"
	"math"
)

const (
	aspectRatio = 16.0 / 9.0

	imageWidth  = 384
	imageHeight = int(imageWidth / aspectRatio)
)

func Run(stdout, stderr io.Writer) error {

	world := newHitters().
		add(newSphere(newPoint(0, 0, -1), 0.5)).
		add(newSphere(newPoint(0, -100.5, -1), 100))

	cam := defaultCamera()

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			fmt.Fprintf(stderr, "\rScanlines remaining: %d", j)
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			c := rayColor(cam.castRay(u, v), world)
			writeColor(stdout, c)
		}
	}
	fmt.Fprintln(stderr, "\nDone.")

	return nil
}

func rayColor(r ray, world hitters) color {
	if hr, ok := world.hit(r, 0, math.MaxFloat64); ok {
		return newColor(0.5*(hr.normal.x+1), 0.5*(hr.normal.y+1), 0.5*(hr.normal.z+1))
	}
	unit := r.direction.normalize()
	x := 0.5 * (unit.y + 1)
	return newColor((1-x)*1.0+x*0.5, (1-x)*1.0+x*0.7, (1-x)*1.0+x*1.0)
}
