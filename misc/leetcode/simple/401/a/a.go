package main

func numberOfChild(n, k int) int {
	t := k % (n - 1)
	if k/(n-1)%2 > 0 {
		return n - t - 1
	}
	return t
}
