package raytracing

import (
	"fmt"
	"io"
	"math"
)

type Color Vector

func NewColor(r, g, b float64) Color {
	return Color(NewVector(r, g, b))
}

func writeColor(w io.Writer, c Color, samples int) {

	clamp := func(x, min, max float64) float64 {
		if x < min {
			return min
		}
		if x > max {
			return max
		}
		return x
	}

	r := int(255 * clamp(math.Sqrt(c.x/float64(samples)), 0, 0.999))
	g := int(255 * clamp(math.Sqrt(c.y/float64(samples)), 0, 0.999))
	b := int(255 * clamp(math.Sqrt(c.z/float64(samples)), 0, 0.999))

	fmt.Fprintf(w, "%d %d %d\n", r, g, b)
}
