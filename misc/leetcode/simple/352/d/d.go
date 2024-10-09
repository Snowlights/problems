package main

func sumImbalanceNumbers(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums {
		vis := make([]int, n+2)
		vis[x] = 1
		cnt := 0
		for j := i + 1; j < n; j++ {
			if x := nums[j]; vis[x] == 0 {
				cnt += 1 - vis[x-1] - vis[x+1]
				vis[x] = 1
			}
			ans += cnt
		}
	}
	return
}
