package main

func isPossibleToSplit(nums []int) bool {
	h := make(map[int]int)
	for _, v := range nums {
		if h[v] == 2 {
			return false
		}
		h[v]++
	}
	return true
}
