package main

import (
	"image"
	"image/color"
	"math"

	"github.com/schollz/progressbar/v3"
	"gonum.org/v1/gonum/mat"
)

func constrain(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}

func normalization(v *mat.VecDense) {
	v.ScaleVec(1/math.Sqrt(mat.Dot(v, v)), v)
}

func Render(width int, height int) (image.Image, error) {
	eyePos := mat.NewVecDense(3, []float64{0, 0, -5})
	spherePos := mat.NewVecDense(3, []float64{0, 0, 5})
	sphereR := 1
	lightPos := mat.NewVecDense(3, []float64{-5, 5, -5})

	pw := mat.NewVecDense(3, []float64{0, 0, 0})

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bar := progressbar.Default(int64(height))
	for y := 0; y < height; y++ {
		bar.Add(1)
		pw.SetVec(1, float64(-2*y)/float64(height-1)+1.0)
		for x := 0; x < width; x++ {
			pw.SetVec(0, float64(2*x)/float64(width-1)-1.0)

			eyeDir := mat.NewVecDense(3, nil)
			eyeDir.SubVec(pw, eyePos)

			tmp := mat.NewVecDense(3, nil)
			tmp.SubVec(eyePos, spherePos)

			a := mat.Dot(eyeDir, eyeDir)
			b := 2 * mat.Dot(eyeDir, tmp)
			c := mat.Dot(tmp, tmp) - float64(sphereR*sphereR)
			d := b*b - 4*a*c

			t := -1.0
			if d == 0 {
				t = -b / (2 * a)
			} else if c > 0 {
				t1 := (-b - math.Sqrt(d)) / (2 * a)
				t2 := (-b + math.Sqrt(d)) / (2 * a)

				if t1 > 0 && t2 > 0 {
					t = math.Min(t1, t2)
				} else {
					t = math.Max(t1, t2)
				}
			}

			col := color.RGBA{100, 149, 237, 255}
			if t > 0 {
				intPos := mat.NewVecDense(3, nil)
				intPos.AddScaledVec(eyePos, t, eyeDir)

				lightDir := mat.NewVecDense(3, nil)
				lightDir.SubVec(lightPos, intPos)
				normalization(lightDir)

				sphereN := mat.NewVecDense(3, nil)
				sphereN.SubVec(intPos, spherePos)
				normalization(sphereN)

				nlDot := constrain(mat.Dot(sphereN, lightDir), 0, 1)
				gray := (uint8)(255 * nlDot)
				col = color.RGBA{gray, gray, gray, 255}
			}

			img.Set(x, y, col)
		}
	}

	return img, nil
}
