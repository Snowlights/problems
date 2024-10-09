package main

func minOperations(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range cnt {
		if c == 1 {
			return -1
		}
		ans += (c + 2) / 3
	}
	return
}
