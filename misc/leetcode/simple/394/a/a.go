package main

import "math/bits"

func numberOfSpecialChars(word string) int {
	a, b := 0, 0
	for _, v := range word {
		if 'a' <= v && v <= 'z' {
			a |= 1 << int(byte(v)-'a')
		} else {
			b |= 1 << int(byte(v)-'A')
		}
	}
	return bits.OnesCount(uint(a & b))
}
