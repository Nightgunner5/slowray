package main

import (
	"image/color"
)

func Point(x, y, z float64) color.RGBA64 {
	if x*x+y*y+z*z < 8*8 {
		return color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff}
	}
	if (x-12)*(x-12)+(y-7)*(y-7)+(z+4)*(z+4) < 3*3 {
		return color.RGBA64{0xffff, 0x0, 0x0, 0xffff}
	}
	return color.RGBA64{0, 0, 0, 0}
}
