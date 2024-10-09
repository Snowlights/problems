package main

func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
		left := x
		for j, c := range cnt[:bias] { // 暴力枚举负数范围 [-50,-1]
			left -= c
			if left <= 0 { // 找到美丽值
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]-- // 离开窗口
	}
	return ans
}
