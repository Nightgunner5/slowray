package main

import (
	"image/color"
)

var fog = color.RGBA64{0x2, 0x4, 0x5, 0x10}

var (
	noise0 = NewNoiseGen(0)
	noise1 = NewNoiseGen(1)
	noise2 = NewNoiseGen(2)
)

func Point(x, y, z float64) color.RGBA64 {
	x /= 5
	y /= 5
	z /= 5

	n0 := noise0.Noise(x, y, z)
	n1 := noise1.Noise(x, y, z)
	n2 := noise2.Noise(x, y, z)

	if n0 >= n1 && n0 >= n2 && n0 >= y {
		return color.RGBA64{0x100, 0, 0, 0x100}
	}
	if n1 >= n0 && n1 >= n2 && n1 >= y {
		return color.RGBA64{0, 0x100, 0, 0x100}
	}
	if n2 >= n0 && n2 >= n1 && n2 >= y {
		return color.RGBA64{0, 0, 0x100, 0x100}
	}
	return color.RGBA64{0, 0, 0, 0}
}
