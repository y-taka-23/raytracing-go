package raytracing

import (
	"fmt"
	"io"
)

const (
	aspectRatio = 16.0 / 9.0

	imageWidth  = 384
	imageHeight = int(imageWidth / aspectRatio)

	viewportWidth  = viewportHeight * aspectRatio
	viewportHeight = 2
	focalLength    = 1
)

func Run(stdout, stderr io.Writer) error {

	horizontal := newVector(viewportWidth, 0, 0)
	vertical := newVector(0, viewportHeight, 0)
	lowerLeftCorner := newPoint(-viewportWidth/2, -viewportHeight/2, -focalLength)

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			fmt.Fprintf(stderr, "\rScanlines remaining: %d", j)
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			dir := origin().to(lowerLeftCorner).
				add(horizontal.mul(u)).
				add(vertical.mul(v))
			c := rayColor(newRay(origin(), dir))
			writeColor(stdout, c)
		}
	}
	fmt.Fprintln(stderr, "\nDone.")

	return nil
}

func rayColor(r ray) color {
	unit := r.direction.normalize()
	t := 0.5 * (unit.y + 1)
	return newColor((1-t)*1.0+t*0.5, (1-t)*1.0+t*0.7, (1-t)*1.0+t*1.0)
}
