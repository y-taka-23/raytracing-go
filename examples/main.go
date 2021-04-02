package main

import (
	"math"

	"github.com/y-taka-23/raytracing-go"
)

func main() {

	width, height := 384, 282

	scene := getScene()
	camera := getCamera(width, height)

	cfg := raytracing.NewRendererConfig(width, height)
	raytracing.NewRenderer(cfg).Render(scene, camera)
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

func getCamera(width, height int) raytracing.Camera {

	cfg := raytracing.NewCameraConfig(
		float64(width)/float64(height),
		raytracing.WithPointOfView(
			raytracing.NewPoint(13, 2, 3),
			raytracing.NewPoint(2, 0.75, 0)),
		raytracing.WithVerticalFOV(math.Pi/9.0),
		raytracing.WithAperture(0.1),
		raytracing.WithFocusDistance(10.0),
	)

	return raytracing.NewCamera(cfg)
}
