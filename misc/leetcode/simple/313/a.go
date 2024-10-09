package simple_313

import "math/bits"

// 1
func commonFactors(a, b int) (ans int) {
	g := gcd(a, b)
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i < g {
				ans++
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// 2
func maxSum(grid [][]int) (ans int) {
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			ans = max(ans, grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1]+grid[i][j]+grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1])
		}
	}
	return
}

// 3
func minimizeXor(num1 int, num2 int) int {

	n1, n2 := bits.OnesCount(uint(num1)), bits.OnesCount(uint(num2))
	for ; n1 < n2; n1++ {
		num1 |= num1 + 1
	}
	for ; n2 < n1; n1-- {
		num1 &= num1 - 1
	}
	return num1
}

// class Solution:
//    def deleteString(self, s: str) -> int:
//        @cache
//        def dfs(s, i):
//            if i == len(s):
//                return 0
//
//            ret = 1
//            span = 1
//            while i + span * 2 <= len(s):
//                if s[i:i+span] == s[i+span:i+span*2]:
//                    ret = max(ret, 1 + dfs(s, i + span))
//                span += 1
//            return ret
//
//
//        ans = dfs(s, 0)
//        dfs.cache_clear()
//        return ans

// 4
func deleteString(s string) int {
	n := len(s)
	lcp := make([][]int, n+1) // lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp[n] = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		lcp[i] = make([]int, n+1)
		for j := n - 1; j > i; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	f := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 1; i+j*2 <= n; j++ {
			if lcp[i][i+j] >= j { // 说明 s[i:i+j] == s[i+j:i+j*2]
				f[i] = max(f[i], f[i+j])
			}
		}
		f[i]++
	}
	return f[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
