package main

func maxNonDecreasingLength(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	nums := [2][]int{nums1, nums2}
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 1
		if nums1[i-1] <= nums[j][i] {
			res = dfs(i-1, 0) + 1
		}
		if nums2[i-1] <= nums[j][i] {
			res = max(res, dfs(i-1, 1)+1)
		}
		*p = res // 记忆化
		return res
	}
	for j := 0; j < 2; j++ {
		for i := 0; i < n; i++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
