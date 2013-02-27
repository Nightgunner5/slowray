package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

var (
	ppi = flag.Int("ppi", 16, "points per image")
	ppp = flag.Int("ppp", 8, "pixels per point")
	spp = flag.Int("spp", 4, "samples per pixel")
	oct = flag.Int("oct", 3, "octaves of noise")

	output  = flag.String("o", "render.png", "output file")
	cpuprof = flag.String("cpuprof", "", "write cpu profile to file")

	basex = flag.Float64("x", 0, "base X coordinate")
	basey = flag.Float64("y", 0, "base Y coordinate")
	basez = flag.Float64("z", -10, "base Z coordinate")
)

const (
	epsilon = 0.01
	farZ    = 200
)

func Pixel(x, y int, img *image.RGBA64, wg *sync.WaitGroup) {
	var r, g, b, a uint64

	for i := 0; i < *spp; i++ {
		for j := 0; j < *spp; j++ {
			X := (float64(x)+float64(i)/float64(*spp))/float64(*ppp) - float64(*ppi)/2
			Y := float64(*ppi)/2 - (float64(y)+float64(j)/float64(*spp))/float64(*ppp)
			c := Ray(X+*basex, Y+*basey, *basez,
				X/float64(*ppi)*epsilon,
				Y/float64(*ppi)*epsilon,
				epsilon,
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

func Worker(ch <-chan [2]int, img *image.RGBA64, wg *sync.WaitGroup) {
	for p := range ch {
		Pixel(p[0], p[1], img, wg)
	}
}

func main() {
	var cpus *int
	if runtime.GOMAXPROCS(0) == 1 {
		cpus = flag.Int("cpus", runtime.NumCPU(), "the number of processor cores to use at any given time")
	} else {
		cpus = flag.Int("cpus", runtime.GOMAXPROCS(0), "the number of processor cores to use at any given time")
	}

	flag.Parse()

	if *cpuprof != "" {
		f, err := os.Create(*cpuprof)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	runtime.GOMAXPROCS(*cpus)

	noise0 = NewNoiseGen(0, *oct)
	noise1 = NewNoiseGen(1, *oct)
	noise2 = NewNoiseGen(2, *oct)

	var wg sync.WaitGroup

	dim := *ppi * *ppp

	img := image.NewRGBA64(image.Rect(0, 0, dim, dim))

	ch := make(chan [2]int, 16)
	for i := 0; i < *cpus; i++ {
		go Worker(ch, img, &wg)
	}

	for x := 0; x < dim; x++ {
		for y := 0; y < dim; y++ {
			wg.Add(1)
			ch <- [2]int{x, y}
		}
	}

	close(ch)

	wg.Wait()

	f, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
