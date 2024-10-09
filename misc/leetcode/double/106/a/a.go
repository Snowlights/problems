package main

import "strconv"

func isFascinating(n int) (ans bool) {

	if n > 333 {
		return false
	}

	mask := 0
	for _, s := range []string{strconv.Itoa(n), strconv.Itoa(n * 2), strconv.Itoa(n * 3)} {
		for _, v := range s {
			mask |= 1 << (v - '0')
		}
	}

	return mask == 1<<10-2
}
