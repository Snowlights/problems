package main

func missingInteger(nums []int) int {
	x, h := nums[0], make(map[int]bool)
	flag := true
	for i, v := range nums {
		if flag && i > 0 && v-nums[i-1] == 1 {
			x += v
		} else if i > 0 {
			flag = false
		}
		h[v] = true
	}
	for h[x] {
		x++
	}
	return x
}
