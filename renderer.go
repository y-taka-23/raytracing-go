package raytracing

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

type Renderer struct {
	imageWriter     io.Writer
	logWriter       io.Writer
	samplesPerPixel int
	maxDepth        int
	imageWidth      int
	imageHeight     int
}

func NewRenderer() Renderer {

	aspectRatio := 4.0 / 3.0
	imageWidth := 384 * 2
	imageHeight := int(float64(imageWidth) / aspectRatio)

	return Renderer{
		imageWriter:     os.Stdout,
		logWriter:       os.Stderr,
		samplesPerPixel: 100,
		maxDepth:        50,
		imageWidth:      imageWidth,
		imageHeight:     imageHeight,
	}
}

func (r Renderer) Render(scene Scene, camera Camera) {

	rand.Seed(time.Now().UnixNano())

	fmt.Fprintln(r.imageWriter, "P3")
	fmt.Fprintf(r.imageWriter, "%d %d\n", r.imageWidth, r.imageHeight)
	fmt.Fprintln(r.imageWriter, 255)

	for j := r.imageHeight - 1; j >= 0; j-- {
		for i := 0; i < r.imageWidth; i++ {
			fmt.Fprintf(r.logWriter, "\rScanlines remaining: %d", j)
			color := NewColor(0, 0, 0)
			for s := 0; s < r.samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(r.imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(r.imageHeight-1)
				c := rayColor(camera.castRay(u, v), scene, r.maxDepth)
				color = NewColor(color.x+c.x, color.y+c.y, color.z+c.z)
			}
			writeColor(r.imageWriter, color, r.samplesPerPixel)
		}
	}
	fmt.Fprintln(r.logWriter, "\nDone.")
}

func rayColor(r ray, scene Scene, depth int) Color {

	if depth <= 0 {
		return NewColor(0, 0, 0)
	}

	hr, ok := scene.hit(r, 0.001, math.MaxFloat64)
	if !ok {
		return backgroundColor(r)
	}

	scattered, ok := hr.material.scatter(hr)
	if !ok {
		return NewColor(0, 0, 0)
	}
	att := hr.material.attenuation()
	c := rayColor(scattered, scene, depth-1)
	return NewColor(att.x*c.x, att.y*c.y, att.z*c.z)
}

func backgroundColor(r ray) Color {
	unit := r.direction.normalize()
	x := 0.5 * (unit.y + 1)
	return NewColor((1-x)*1.0+x*0.5, (1-x)*1.0+x*0.7, (1-x)*1.0+x*1.0)
}
