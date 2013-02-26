package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

const (
	ppp  = 4 // pixels per point
	spp  = 4 // samples per pixel
	spp2 = spp * spp
)

func Pixel(x, y int, img *image.RGBA64, wg *sync.WaitGroup) {
	var r, g, b, a uint64

	for i := 0; i < spp; i++ {
		for j := 0; j < spp; j++ {
			c := Ray((float64(x)+float64(i)/spp)/ppp, (float64(y)+float64(j)/spp)/ppp, -10, 0, 0, 0.001, 100)
			r += uint64(c.R)
			g += uint64(c.G)
			b += uint64(c.B)
			a += uint64(c.A)
		}
	}
	img.SetRGBA64(x, y, color.RGBA64{uint16(r / spp2), uint16(g / spp2), uint16(b / spp2), uint16(a / spp2)})
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
