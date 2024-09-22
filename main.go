package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func make_image(pattern func(uint8, uint8) uint8) (*image.RGBA, error) {
	width := 250
	height := 250
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			value := pattern(uint8(x), uint8(y))
			img.Set(x, y, color.RGBA{0, value, 0, 255})
		}
	}
	return img, nil
}

func main() {
	patterns := []func(uint8, uint8) uint8{
		func(x, y uint8) uint8 { return x * y },
		func(x, y uint8) uint8 { return (x * y) / 2 },
		func(x, y uint8) uint8 { return x + y },
		func(x, y uint8) uint8 { return (x + y) / 2 },
		func(x, y uint8) uint8 { return x ^ y },
	}
	if _, err := os.Stat("img"); os.IsNotExist(err) {
		err := os.Mkdir("img", 0750)
		if err != nil {
			log.Fatal(err)
		}
	}
	for index, pattern := range patterns {

		img, err := make_image(pattern)
		if err != nil {
			log.Fatal(err)
		}

		name := fmt.Sprintf("./img/image_%v.png", index)
		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		if err := png.Encode(f, img); err != nil {
			f.Close()
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
