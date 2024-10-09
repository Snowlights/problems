package main

import "math/bits"

func makeTheIntegerZero(x, y int) int {

	for k := 1; k <= x-k*y; k++ {
		if k >= bits.OnesCount(uint(x-k*y)) {
			return k
		}
	}

	return -1
}
