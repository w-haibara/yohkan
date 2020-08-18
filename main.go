package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/schollz/progressbar/v3"
)

func main() {
	var img_w, img_h int = 256, 256
	bar := progressbar.Default(int64(img_h))

	m := image.NewRGBA(image.Rect(0, 0, img_w, img_h))
	for y := 0; y < img_h; y++ {
		bar.Add(1)
		for x := 0; x < img_w; x++ {
			var (
				r, g, b, a uint8
			)

			r = uint8(x % 255)
			g = uint8(y % 255)
			b = 0
			a = 255

			m.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, m)
}
