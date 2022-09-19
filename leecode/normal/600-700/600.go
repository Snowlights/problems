package _00_700

import "sort"

// 611
func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}

// 621
func leastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for _, v := range tasks {
		cnt[v-'A']++
	}
	sort.Ints(cnt)
	minTime := (n+1)*(cnt[25]-1) + 1
	for i := 0; i < 25; i++ {
		if cnt[i] == cnt[25] {
			minTime++
		}
	}
	return max(minTime, len(tasks))
}
