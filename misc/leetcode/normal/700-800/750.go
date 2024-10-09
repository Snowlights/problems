package _00_800

// 763
func partitionLabels(s string) []int {

	lastPos, ans := [26]int{}, []int{}
	for i, v := range s {
		lastPos[v-'a'] = i
	}
	start, end := 0, 0
	for i, v := range s {
		if lastPos[v-'a'] > end {
			end = lastPos[v-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}

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
