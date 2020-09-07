package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"yohkan/renderer"
)

func encodeToPNG(img image.Image, path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	const aspectRatio = 1
	const imgWidth = 5120
	const imgHeight = (int)(imgWidth / aspectRatio)

	img, err := renderer.Render(imgWidth, imgHeight)
	if err != nil {
		log.Fatal(err)
	}

	encodeToPNG(img, "out.png")
}
