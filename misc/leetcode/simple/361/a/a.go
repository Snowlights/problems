package main

import "strconv"

func countSymmetricIntegers(low int, high int) (ans int) {
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n&1 == 1 {
			continue
		}
		a, b := 0, 0
		for _, v := range s[:n/2] {
			a += int(v - '0')
		}
		for _, v := range s[n/2:] {
			b += int(v - '0')
		}
		if a == b {
			ans++
		}
	}

	return
}
