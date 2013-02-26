package main

import (
	"image"
	"image/png"
	"os"
	"sync"
)

const ppp = 4 // pixels per point

func Pixel(x, y int, img *image.RGBA64, wg *sync.WaitGroup) {
	img.SetRGBA64(x, y, Ray(float64(x)/ppp, float64(y)/ppp, -10, 0, 0, 0.001, 20))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	const dim = 16 * ppp

	img := image.NewRGBA64(image.Rect(0, 0, dim, dim))

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			wg.Add(1)
			go Pixel(x, y, img, &wg)
		}
	}

	wg.Wait()

	f, err := os.Create("render.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
