package _500_1600

// 1567
func getMaxLen(nums []int) (ans int) {
	pos, neg := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos++
			if neg > 0 {
				neg++
			}
		} else if num == 0 {
			pos, neg = 0, 0
		} else {
			if neg > 0 {
				pos, neg = neg+1, pos+1
			} else {
				pos, neg = 0, pos+1
			}
		}
		ans = max(ans, pos)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
