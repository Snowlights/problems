package simple_305

import "fmt"

// 1
func arithmeticTriplets(nums []int, diff int) int {

	h := make(map[int]bool, len(nums))
	for _, n := range nums {
		h[n] = true
	}
	ans := 0
	for _, v := range nums {
		if h[v-diff] && h[v+diff] {
			ans++
		}
	}
	return ans
}

// 2
func reachableNodes(n int, edges [][]int, restricted []int) int {

	m := make(map[int][]int)
	restMap := make(map[int]bool)
	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}
	for _, v := range restricted {
		restMap[v] = true
	}

	reached := make(map[int]bool)
	var dfs func(int)
	dfs = func(n int) {
		if restMap[n] || reached[n] {
			return
		}
		reached[n] = true
		for _, v := range m[n] {
			dfs(v)
		}
	}
	dfs(0)
	fmt.Println(reached)
	return len(reached)
}

// 3
func validPartition(nums []int) bool {

	dp := make([]bool, len(nums)+1)
	dp[0] = true
	for i, v := range nums {
		if (i > 0 && dp[i-1] && nums[i-1] == v) ||
			(i > 1 && dp[i-2] && (v == nums[i-1] && v == nums[i-2] ||
				v == nums[i-1]+1 && v == nums[i-2]+2)) {
			dp[i+1] = true
		}
	}
	return dp[len(nums)]
}

// 4

func longestIdealString(s string, k int) (ans int) {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		for j := max(c-k, 0); j <= min(c+k, 25); j++ {
			f[c] = max(f[c], f[j])
		}
		f[c]++
	}
	for _, v := range f {
		ans = max(ans, v)
	}
	return
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
