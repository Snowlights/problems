package main

func getGoodIndices(variables [][]int, target int) []int {
	ans := make([]int, 0, len(variables))
	pow := func(x, n, m int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % m
			}
			x = x * x % m
		}
		return res
	}
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return ans
}
