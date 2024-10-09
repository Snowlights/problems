package interview

import (
	"fmt"
)

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

// 面试题 05-04

func findClosedNumbers(num int) []int {
	ret := []int{-1, -1}
	s := fmt.Sprintf("0%b", num)
	p, count := 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '1' && s[i-1] == '0' && ret[0] == -1 {
			ret[0] = (num>>p+1)<<p + 1<<count - 1
		}
		if s[i] == '0' && s[i-1] == '1' && ret[1] == -1 {
			ret[1] = (num>>p-1)<<p + 1<<p - 1<<(p-count)
		}
		if s[i] == '1' {
			count++
		}
		p++
	}
	if ret[0] > 2147483647 || ret[0] < 0 {
		ret[0] = -1
	}

	if ret[1] > 2147483647 || ret[1] < 0 {
		ret[1] = -1
	}
	return ret
}
