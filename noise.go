package main

import (
	"math"
	"math/rand"
)

type NoiseGen [256]uint8

func NewNoiseGen(seed int64) *NoiseGen {
	var n NoiseGen
	r := rand.New(rand.NewSource(seed))

	perm := r.Perm(len(n))

	for i := range n {
		n[i] = uint8(perm[i])
	}

	return &n
}

func (n *NoiseGen) Noise(x, y, z float64) float64 {
	mul := 1.0
	cur := 0.0
	max := 0.0

	for i := 0; i < 3; i++ {
		cur += n.noise(x/mul, y/mul, z/mul) * mul
		max += mul
		mul *= 0.25
	}

	return cur / max
}

// Adapted from http://mrl.nyu.edu/~perlin/noise/
func (n *NoiseGen) noise(x, y, z float64) float64 {
	// Find unit cube that contains point.
	X := uint8(math.Floor(x))
	Y := uint8(math.Floor(y))
	Z := uint8(math.Floor(z))

	// Find relative X, Y, Z of point in cube.
	x -= math.Floor(x)
	y -= math.Floor(y)
	z -= math.Floor(z)

	// Compute fade curves for each of X, Y, Z.
	u, v, w := n.fade(x), n.fade(y), n.fade(z)

	// Hash coordinates of the 8 cube corners,
	var (
		A  = (*n)[X] + Y
		AA = (*n)[A] + Z
		AB = (*n)[A+1] + Z
		B  = (*n)[X+1] + Y
		BA = (*n)[B] + Z
		BB = (*n)[B+1] + Z
	)

	// and add blended results from 8 corners of cube.
	return n.lerp(w, n.lerp(v, n.lerp(u,
		n.grad((*n)[AA], x, y, z),
		n.grad((*n)[BA], x-1, y, z)),
		n.lerp(u,
			n.grad((*n)[AB], x, y-1, z),
			n.grad((*n)[BB], x-1, y-1, z))),
		n.lerp(v, n.lerp(u,
			n.grad((*n)[AA+1], x, y, z-1),
			n.grad((*n)[BA+1], x-1, y, z-1)),
			n.lerp(u,
				n.grad((*n)[AB+1], x, y-1, z-1),
				n.grad((*n)[BB+1], x-1, y-1, z-1))))
}
func (*NoiseGen) fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}
func (*NoiseGen) lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}
func (*NoiseGen) grad(hash uint8, x, y, z float64) float64 {
	// Convert low four bits of hash code into twelve gradient directions.
	h := hash & 15
	u := y
	if h < 8 {
		u = x
	}
	v := z
	if h < 4 {
		v = y
	} else if h == 12 || h == 14 {
		v = x
	}

	s1 := 1 - float64((h&1)<<1)
	s2 := 1 - float64(h&2)
	return u*s1 + v*s2
}
