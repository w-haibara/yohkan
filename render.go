package main

import (
	"image"
	"image/color"

	"github.com/schollz/progressbar/v3"
	"gonum.org/v1/gonum/mat"
)

func Render(width int, height int) (image.Image, error) {
	eyePos := mat.NewVecDense(3, []float64{0, 0, -5})

	spherePos := mat.NewVecDense(3, []float64{0, 0, 5})
	sphereR := 1

	pw := mat.NewVecDense(3, []float64{0, 0, 0})

	m := image.NewRGBA(image.Rect(0, 0, width, height))
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

			col := color.RGBA{0, 0, 0, 255}
			if d >= 0 {
				col = color.RGBA{255, 0, 0, 255}
			} else {
				col = color.RGBA{0, 0, 255, 255}
			}

			m.Set(x, y, col)
		}
	}

	return m, nil
}
