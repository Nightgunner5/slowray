package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
)

var (
	ppi = flag.Int("ppi", 16, "points per image")
	ppp = flag.Int("ppp", 8, "pixels per point")
	spp = flag.Int("spp", 4, "samples per pixel")
)

const (
	epsilon = 0.01
	farZ    = 100
)

func Pixel(x, y int, img *image.RGBA64, wg *sync.WaitGroup) {
	var r, g, b, a uint64

	for i := 0; i < *spp; i++ {
		for j := 0; j < *spp; j++ {
			X := (float64(x)+float64(i)/float64(*spp))/float64(*ppp) - float64(*ppi)/2
			Y := float64(*ppi)/2 - (float64(y)+float64(j)/float64(*spp))/float64(*ppp)
			c := Ray(X, Y, -10,
				X/float64(*ppi)*epsilon,
				Y/float64(*ppi)*epsilon,
				1*epsilon,
				farZ)

			r += uint64(c.R)
			g += uint64(c.G)
			b += uint64(c.B)
			a += uint64(c.A)
		}
	}
	img.SetRGBA64(x, y, color.RGBA64{
		uint16(r / uint64(*spp) / uint64(*spp)),
		uint16(g / uint64(*spp) / uint64(*spp)),
		uint16(b / uint64(*spp) / uint64(*spp)),
		uint16(a / uint64(*spp) / uint64(*spp)),
	})
	wg.Done()
}

func main() {
	var cpus *int
	if runtime.GOMAXPROCS(0) == 1 {
		cpus = flag.Int("cpus", runtime.NumCPU(), "the number of processor cores to use at any given time")
	} else {
		cpus = flag.Int("cpus", runtime.GOMAXPROCS(0), "the number of processor cores to use at any given time")
	}

	flag.Parse()

	runtime.GOMAXPROCS(*cpus)

	var wg sync.WaitGroup

	dim := *ppi * *ppp

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
