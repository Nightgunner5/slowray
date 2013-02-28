package main

import (
	"math"
	"math/rand"
)

type noiseGen [256]uint8
type NoiseGen []noiseGen

func NewNoiseGen(seed int64, octaves int) NoiseGen {
	n := make(NoiseGen, octaves)
	r := rand.New(rand.NewSource(seed))

	for i := range n {
		perm := r.Perm(len(n[i]))

		for j := range n[i] {
			n[i][j] = uint8(perm[j])
		}
	}

	return n
}

func (n NoiseGen) Noise(x, y, z float64) float64 {
	mul := 1.0
	cur := 0.0
	max := 0.0

	for i := range n {
		cur += n[i].noise(x/mul, y/mul, z/mul) * mul
		max += mul
		mul *= 0.25
	}

	return cur / max
}

func noisegen_floor(f float64) (mod uint8, frac float64) {
	fl := math.Floor(f)
	return uint8(fl), f - fl
}

// Adapted from http://mrl.nyu.edu/~perlin/noise/
func (n noiseGen) noise(x, y, z float64) float64 {
	// Find unit cube that contains point.
	// Find relative X, Y, Z of point in cube.
	X, x := noisegen_floor(x)
	Y, y := noisegen_floor(y)
	Z, z := noisegen_floor(z)

	// Compute fade curves for each of X, Y, Z.
	u, v, w := noisegen_fade(x), noisegen_fade(y), noisegen_fade(z)

	// Hash coordinates of the 8 cube corners,
	var (
		A  = n[X] + Y
		AA = n[A] + Z
		AB = n[A+1] + Z
		B  = n[X+1] + Y
		BA = n[B] + Z
		BB = n[B+1] + Z
	)

	// and add blended results from 8 corners of cube.
	return noisegen_lerp(w, noisegen_lerp(v, noisegen_lerp(u,
		noisegen_grad(n[AA], x, y, z),
		noisegen_grad(n[BA], x-1, y, z)),
		noisegen_lerp(u,
			noisegen_grad(n[AB], x, y-1, z),
			noisegen_grad(n[BB], x-1, y-1, z))),
		noisegen_lerp(v, noisegen_lerp(u,
			noisegen_grad(n[AA+1], x, y, z-1),
			noisegen_grad(n[BA+1], x-1, y, z-1)),
			noisegen_lerp(u,
				noisegen_grad(n[AB+1], x, y-1, z-1),
				noisegen_grad(n[BB+1], x-1, y-1, z-1))))
}
func noisegen_fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}
func noisegen_lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
func noisegen_grad(hash uint8, x, y, z float64) float64 {
	// Convert low four bits of hash code into twelve gradient directions.
	switch uint(hash & 15) {
	case 0:
		return x + y
	case 1:
		return -x + y
	case 2:
		return x - y
	case 3:
		return -x - y
	case 4:
		return x + z
	case 5:
		return -x + z
	case 6:
		return x - z
	case 7:
		return -x - z
	case 8:
		return y + z
	case 9:
		return -y + z
	case 10:
		return y - z
	case 11:
		return -y - z
	case 12:
		return y + x
	case 13:
		return -y + z
	case 14:
		return y - x
	case 15:
		return -y - z
	}
	panic("unreachable")
}
