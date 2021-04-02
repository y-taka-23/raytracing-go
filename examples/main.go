package main

import (
	"math"
	"os"

	"github.com/y-taka-23/raytracing-go"
)

func main() {
	scene := getScene()
	camera := getCamera()
	raytracing.Render(os.Stdout, os.Stderr, scene, camera)
}

func getScene() raytracing.Scene {

	scene := raytracing.NewScene()

	scene.Add(raytracing.NewSphere(
		raytracing.NewPoint(0, -1000, 0), 1000,
		raytracing.NewLambertian(raytracing.NewColor(0.5, 0.5, 0.5))))
	scene.Add(raytracing.NewSphere(
		raytracing.NewPoint(0, 1, 0), 1,
		raytracing.NewDielectric(1.5)))
	scene.Add(raytracing.NewSphere(
		raytracing.NewPoint(-4, 1, 0), 1,
		raytracing.NewLambertian(raytracing.NewColor(0.4, 0.2, 0.1))))
	scene.Add(raytracing.NewSphere(
		raytracing.NewPoint(4, 1, 0), 1,
		raytracing.NewMetal(raytracing.NewColor(0.7, 0.6, 0.5), 0)))

	return *scene
}

func getCamera() raytracing.Camera {

	lookFrom := raytracing.NewPoint(13, 2, 3)
	lookAt := raytracing.NewPoint(2, 0.75, 0)
	viewUp := raytracing.NewVector(0, 1, 0)
	aspectRatio := 4.0 / 3.0
	aperture := 0.1
	distToFocus := 10.0

	return raytracing.NewCamera(lookFrom, lookAt, viewUp,
		math.Pi/9.0, aspectRatio, aperture, distToFocus)
}
