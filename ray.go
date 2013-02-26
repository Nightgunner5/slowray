package main

import (
	"image/color"
	"math"
)

func Ray(x, y, z, dx, dy, dz, maxDistance float64) color.RGBA64 {
	var result color.RGBA64

	dist := 0.0
	ddist := math.Sqrt(dx*dx + dy*dy + dz*dz)

	for result.A != 0xffff && dist < maxDistance {
		c := Point(x, y, z)
		sr := uint64(result.R) * 0x10001
		sg := uint64(result.G) * 0x10001
		sb := uint64(result.B) * 0x10001
		sa := uint64(result.A) * 0x10001

		dr := uint64(c.R)
		dg := uint64(c.G)
		db := uint64(c.B)
		da := uint64(c.A)

		a := (0xffffffff - sa) * 0x10001

		result.R = uint16((dr*a/0xffffffff + sr) >> 16)
		result.G = uint16((dg*a/0xffffffff + sg) >> 16)
		result.B = uint16((db*a/0xffffffff + sb) >> 16)
		result.A = uint16((da*a/0xffffffff + sa) >> 16)

		x += dx
		y += dy
		z += dz
		dist += ddist
	}

	return result
}
