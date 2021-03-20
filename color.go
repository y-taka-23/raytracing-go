package raytracing

import (
	"fmt"
	"io"
)

type color vector

func newColor(r, g, b float64) color {
	return color(newVector(r, g, b))
}

func writeColor(w io.Writer, c color, samples int) {

	clamp := func(x, min, max float64) float64 {
		if x < min {
			return min
		}
		if x > max {
			return max
		}
		return x
	}

	r := int(255 * clamp(c.x/float64(samples), 0, 0.999))
	g := int(255 * clamp(c.y/float64(samples), 0, 0.999))
	b := int(255 * clamp(c.z/float64(samples), 0, 0.999))
	fmt.Fprintf(w, "%d %d %d\n", r, g, b)
}
