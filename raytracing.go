package raytracing

import (
	"fmt"
	"io"
)

const (
	imageWidth  = 256
	imageHeight = 256
)

func Run(stdout, stderr io.Writer) error {

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			fmt.Fprintf(stderr, "\rScanlines remaining: %d", j)

			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(r * 255.999)
			ig := int(g * 255.999)
			ib := int(b * 255.999)

			fmt.Fprintf(stdout, "%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprintln(stderr, "\nDone.")

	return nil
}
