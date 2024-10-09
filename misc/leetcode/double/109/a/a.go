package main

func isGood(nums []int) bool {
	n := len(nums) - 1
	vis := make([]bool, n+1)
	for _, v := range nums {
		if v > n || v < n && vis[v] {
			return false
		}
		vis[v] = true
	}
	return true
}
