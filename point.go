package main

import (
	"image/color"
)

func Point(x, y, z float64) color.RGBA64 {
	if x*x+y*y+z*z < 8*8 {
		return color.RGBA64{0x10, 0x10, 0x10, 0x10}
	}
	return color.RGBA64{0, 0, 0, 0}
}
