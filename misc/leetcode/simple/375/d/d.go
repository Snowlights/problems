package main

func numberOfGoodPartitions(nums []int) int {
	h, ans, r := make(map[int]int), 0, -1
	for i, v := range nums {
		h[v] = i
	}
	const mod int = 1e9 + 7
	pow := func(x, n int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	for i, v := range nums {
		r = max(r, h[v])
		if r == i {
			ans++
		}
	}
	return pow(2, ans-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
