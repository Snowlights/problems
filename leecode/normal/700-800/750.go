package _00_800

// 769
func maxChunksToSorted(arr []int) int {
	mx, ans := arr[0], 0
	for i, v := range arr {
		mx = int(max(int64(mx), int64(v)))
		if mx == i {
			ans++
		}
	}
	return ans
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
