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
		result = Add(result, Point(x, y, z))
		result = Add(result, fog)

		x += dx
		y += dy
		z += dz
		dist += ddist
	}

	return result
}

func Add(top, bottom color.RGBA64) color.RGBA64 {
	sr := uint64(top.R) * 0x10001
	sg := uint64(top.G) * 0x10001
	sb := uint64(top.B) * 0x10001
	sa := uint64(top.A) * 0x10001

	dr := uint64(bottom.R)
	dg := uint64(bottom.G)
	db := uint64(bottom.B)
	da := uint64(bottom.A)

	a := (0xffffffff - sa) * 0x10001

	return color.RGBA64{
		R: uint16((dr*a/0xffffffff + sr) >> 16),
		G: uint16((dg*a/0xffffffff + sg) >> 16),
		B: uint16((db*a/0xffffffff + sb) >> 16),
		A: uint16((da*a/0xffffffff + sa) >> 16),
	}
}
