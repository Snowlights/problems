package main

func canSplitArray(a []int, m int) (ans bool) {
	if len(a) <= 2 {
		return true
	}
	for i := 1; i < len(a); i++ {
		if a[i]+a[i-1] >= m {
			return true
		}
	}
	return
}
