package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func encodeToPNG(m image.Image, path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, m)
}

func main() {
	const aspectRatio = 1
	const imgWidth = 512
	const imgHeight = (int)(imgWidth / aspectRatio)

	m, err := Render(imgWidth, imgHeight)
	if err != nil {
		log.Fatal(err)
	}

	encodeToPNG(m, "out.png")
}
