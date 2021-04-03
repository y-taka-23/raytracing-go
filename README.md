# Ray Tracing in Go

![cover image](images/cover.png)

A Go implementation of the book [Ray Tracing in One Weekend](https://raytracing.github.io/). The repository provides a library to describe and render your own scenes. For more detail, see [`examples/main.go`](examples/main.go).

## Getting Started

```shell
git clone git@github.com:y-taka-23/raytracing-go.git
cd raytracing-go
make examples
./bin/example > example.ppm
open example.ppm
```

## Materials

### Lambertian

|color|result|
|:----:|:----:|
|(0.8, 0.1, 0.1)|![red lambertian sphere](images/lambertian_red.png)|
|(1.0, 1.0, 1.0)|![white lambertian sphere](images/lambertian_white.png)|
|(0.0, 0.0, 0.0)|![black lambertian sphere](images/lambertian_black.png)|

### Metalic

|fuzziness|result|
|:----:|:----:|
|0.0|![metalic sphere of fuzziness 0.0](images/metal_f00.png)|
|0.15|![metalic sphere of fuzziness 0.15](images/metal_f15.png)|
|0.3|![metalic sphere of fuzziness 0.3](images/metal_f30.png)|

### Dielectric

|refractive index|result|
|:----:|:----:|
|1.0|![dielectric sphere of refractive index 1.0](images/dielectric_i100.png)|
|1.5|![dielectric sphere of refractive index 1.5](images/dielectric_i150.png)|
|2.0|![dielectric sphere of refractive index 2.0](images/dielectric_i200.png)|


## Camera Setting

### Angle of View

|virtical angle (degree)|result|
|:----:|:----:|
|90|![result of the vertical angle in 90 degree](images/fov_90.png)|
|60|![result of the vertical angle in 60 degree](images/fov_60.png)|
|30|![result of the vertical angle in 30 degree](images/fov_30.png)|

### Aperture

|aperture|result|
|:----:|:----:|
|0.0|![result of the aperture 0.0](images/aperture_000.png)|
|0.5|![result of the aperture 0.5](images/aperture_050.png)|
|1.0|![result of the aperture 1.0](images/aperture_100.png)|

### Depth of Field

|focus distance|result|
|:----:|:----:|
|6|![result of the depth of field 6](images/dof_06.png)|
|9|![result of the depth of field 9](images/dof_09.png)|
|12|![result of the depth of field 12](images/dof_12.png)|

## Reference

* Shirley, P. (2016). _Ray Tracing in One Weekend_ ([online edition available](https://raytracing.github.io/))
