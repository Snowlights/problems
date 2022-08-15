package _200_1300

// 1281
func groupThePeople(groupSizes []int) [][]int {

	h := make(map[int][]int)
	for i, v := range groupSizes {
		h[v] = append(h[v], i)
	}
	ans := make([][]int, 0)
	for k, v := range h {
		for i := 0; i < len(v); i += k {
			ans = append(ans, v[i:i+k])
		}
	}
	return ans
}
