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

func Render(stdout, stderr io.Writer, scene Scene, camera Camera) {

	rand.Seed(time.Now().UnixNano())

	fmt.Fprintln(stdout, "P3")
	fmt.Fprintf(stdout, "%d %d\n", imageWidth, imageHeight)
	fmt.Fprintln(stdout, 255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			fmt.Fprintf(stderr, "\rScanlines remaining: %d", j)
			color := NewColor(0, 0, 0)
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				c := rayColor(camera.castRay(u, v), scene, maxDepth)
				color = NewColor(color.x+c.x, color.y+c.y, color.z+c.z)
			}
			writeColor(stdout, color, samplesPerPixel)
		}
	}
	fmt.Fprintln(stderr, "\nDone.")
}

func rayColor(r ray, world Scene, depth int) Color {

	if depth <= 0 {
		return NewColor(0, 0, 0)
	}

	hr, ok := world.hit(r, 0.001, math.MaxFloat64)
	if !ok {
		unit := r.direction.normalize()
		x := 0.5 * (unit.y + 1)
		return NewColor((1-x)*1.0+x*0.5, (1-x)*1.0+x*0.7, (1-x)*1.0+x*1.0)
	}

	scattered, ok := hr.material.scatter(hr)
	if !ok {
		return NewColor(0, 0, 0)
	}
	att := hr.material.attenuation()
	c := rayColor(scattered, world, depth-1)
	return NewColor(att.x*c.x, att.y*c.y, att.z*c.z)
}
