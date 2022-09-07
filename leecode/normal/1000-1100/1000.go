package _000_1100

// 1014
func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func maxScoreSightseeingPair(values []int) int {

	preM, ans := int64(0), int64(0)
	for i, v := range values {
		if i == 0 {
			preM = int64(v + i)
			continue
		}
		ans = max(ans, int64(preM+int64(v-i)))
		preM = max(preM, int64(v+i))
	}

	return int(ans)
}
