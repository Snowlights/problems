package interview

// 面试题 05-01
func insertBits(N int, M int, i int, j int) int {
	// 分别截断高位 替换位 低位
	H := (((1 << (31 - j)) - 1) << (j + 1)) & N
	M = M << i
	L := N & ((1 << i) - 1)
	return H | M | L
}

// 面试题 05-02
func reverseBits(num int) int {

	a := make([]int, 0)
	for i := 0; i < 32; i++ {
		a = append(a, num>>i&1)
	}

	cal := func() int {
		res, cur := 0, 0
		for _, v := range a {
			if v == 1 {
				cur++
			} else {
				res = max(res, cur)
				cur = 0
			}
		}
		return max(res, cur)
	}
	ans := cal()
	for i, v := range a {
		if v == 0 {
			a[i] = 1
			ans = max(ans, cal())
			a[i] = 0
		}
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
