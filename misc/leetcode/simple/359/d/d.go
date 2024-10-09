package main

func longestEqualSubarray(nums []int, k int) (ans int) {
	pos := make([][]int, len(nums)+1)
	for i, x := range nums {
		pos[x] = append(pos[x], i-len(pos[x]))
	}
	for _, ps := range pos {
		left := 0
		for right, p := range ps {
			for p-ps[left] > k { // 要删除的数太多了
				left++
			}
			ans = max(ans, right-left+1)
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
