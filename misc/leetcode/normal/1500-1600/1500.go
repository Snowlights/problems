package _500_1600

import (
	"fmt"
	"sort"
)

// 1508
func rangeSum(nums []int, n int, left int, right int) int {
	ans := make([]int64, 0, (n+1)*n/2)
	mod := int64(1e9 + 7)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			ans = append(ans, prefix[j]-prefix[i])
		}
	}
	sort.Slice(ans, func(i int, j int) bool {
		return ans[i] < ans[j]
	})
	res := int64(0)
	for i := left - 1; i < right; i++ {
		res += ans[i]
		res %= mod
	}
	return int(res)
}

// 1545
func findKthBit(n int, k int) byte {
	s := []byte{'0'}

	reserve := func(b []byte) []byte {
		ans := make([]byte, 0, len(b))
		for _, v := range b {
			switch v {
			case '1':
				ans = append(ans, '0')
			case '0':
				ans = append(ans, '1')
			}
		}
		s, e := 0, len(ans)-1
		for s < e {
			ans[s], ans[e] = ans[e], ans[s]
			s++
			e--
		}
		return ans
	}

	for i := 1; i < n; i++ {
		s = append(s, append([]byte{'1'}, reserve(s)...)...)
		fmt.Println(string(s))
	}
	return s[k]
}
