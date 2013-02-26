package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
)

const (
	ppi  = 16 // points per image
	ppp  = 32 // pixels per point
	spp  = 4  // samples per pixel
	spp2 = spp * spp

	epsilon = 0.01
	nearZ   = 1
	farZ    = 100
)

func Pixel(x, y int, img *image.RGBA64, wg *sync.WaitGroup) {
	var r, g, b, a uint64

	for i := 0; i < spp; i++ {
		for j := 0; j < spp; j++ {
			X := (float64(x)+float64(i)/spp)/ppp - ppi/2
			Y := ppi/2 - (float64(y)+float64(j)/spp)/ppp
			c := Ray(X, Y, -10,
				X/ppi*epsilon,
				Y/ppi*epsilon,
				1*epsilon,
				farZ)

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
	if runtime.GOMAXPROCS(0) == 1 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var wg sync.WaitGroup

	const dim = ppi * ppp

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
