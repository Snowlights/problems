package main

func minimumOperations(nums, target []int) int64 {
	s := target[0] - nums[0]
	ans := abs(s)
	for i := 1; i < len(nums); i++ {
		k := (target[i] - target[i-1]) - (nums[i] - nums[i-1])
		if k > 0 {
			if s >= 0 {
				ans += k
			} else {
				ans += max(k+s, 0)
			}
		} else {
			if s <= 0 {
				ans -= k
			} else {
				ans -= min(k+s, 0)
			}
		}
		s += k
	}
	return int64(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
