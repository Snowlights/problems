package simple_333

import (
	"bytes"
	"sort"
)

// 1
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	h := make(map[int][]int)
	for _, v := range nums1 {
		h[v[0]] = v
	}
	for _, v := range nums2 {
		if _, ok := h[v[0]]; ok {
			h[v[0]][1] += v[1]
		} else {
			h[v[0]] = v
		}
	}
	ans := make([][]int, 0, len(h))
	for _, v := range h {
		ans = append(ans, v)
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

// 2
func minOperations(n int) int {
	lowbit := func(x int) int {
		return x & (-x)
	}

	var dfs func(x int) int
	dfs = func(x int) int {
		if x&(x-1) == 0 {
			return 1
		}
		return 1 + min(dfs(x+lowbit(x)), dfs(x-lowbit(x)))
	}

	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 3
func squareFreeSubsets(nums []int) int {
	mod := int(1e9 + 7)
	// m, Mask := 1<<10, []int{-1, 0, 1, 2, -1, 4, 3, 8, -1, -1, 5, 16, -1, 32, 9, 6, -1, 64, -1, 128, -1, 10, 17, 256, -1, -1, 33, -1, -1, 512, 7}
	m := 1 << 10
	p := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	Mask := [31]int{}
	for i := 2; i < 31; i++ {
		for j, p := range p {
			if i%p == 0 {
				if i%(p*p) == 0 {
					Mask[i] = -1
					break
				}
				Mask[i] |= 1 << j
			}
		}
	}
	f := make([]int, m)
	f[0] = 1
	for i := 1; i < len(nums)+1; i++ {
		if mask := Mask[nums[i-1]]; mask >= 0 {
			for j := m - 1; j >= 0; j-- {
				if mask&j == mask {
					f[j] = (f[j] + f[j^mask]) % mod
				}
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
		ans %= mod
	}
	return (ans - 1 + mod) % mod
}

// 4
func findTheString(lcp [][]int) string {
	n := len(lcp)
	for i, row := range lcp {
		if row[i] != n-i { // 和自己的 LCP 必然等于整个后缀长度
			return ""
		}
		for j := i + 1; j < n; j++ {
			if row[j] != lcp[j][i] || row[j] > n-max(i, j) { // 不对称 or 超过整个后缀长度
				return ""
			}
		}
	}

	s := make([]byte, n)
	for c := byte('a'); c <= 'z'; c++ {
		i := bytes.IndexByte(s, 0)
		if i < 0 {
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				s[j] = c
			}
		}
	}
	if bytes.IndexByte(s, 0) >= 0 { // 还有没构造完的
		return ""
	}

	// 直接在原数组上验证（DP 算 LCP）
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == n-1 || j == n-1 {
					if lcp[i][j] != 1 {
						return ""
					}
				} else if lcp[i][j] != lcp[i+1][j+1]+1 {
					return ""
				}
			} else if lcp[i][j] > 0 { // 不相等应该是 0
				return ""
			}
		}
	}
	return string(s)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
