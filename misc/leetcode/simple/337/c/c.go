package main

func beautifulSubsets(nums []int, k int) int {
	ans := -1                    // 去掉空集
	cnt := make([]int, 1001+k*2) // 用数组实现比哈希表更快
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1)       // 不选
		x := nums[i] + k // 避免负数下标
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++ // 选
			dfs(i + 1)
			cnt[x]-- // 恢复现场
		}
	}
	dfs(0)
	return ans
}
