package raytracing

import (
	"fmt"
)

const (
	imageWidth  = 256
	imageHeight = 256
)

func Run() error {

	fmt.Println("P3")
	fmt.Printf("%d %d\n", imageWidth, imageHeight)
	fmt.Println(255)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {

			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(r * 255.999)
			ig := int(g * 255.999)
			ib := int(b * 255.999)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}

	return nil
}
