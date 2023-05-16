package main

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}
