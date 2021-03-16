package raytracing

import (
	"fmt"
	"io"
)

type color vector

func newColor(r, g, b float64) color {
	return color(newVector(r, g, b))
}

func writeColor(w io.Writer, c color) {
	ir := int(c.x * 255.999)
	ig := int(c.y * 255.999)
	ib := int(c.z * 255.999)
	fmt.Fprintf(w, "%d %d %d\n", ir, ig, ib)
}
