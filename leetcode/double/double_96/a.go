package double_96

// 1
func getCommon(nums1 []int, nums2 []int) int {
	h := make(map[int]bool)
	for _, v := range nums1 {
		h[v] = true
	}
	for _, v := range nums2 {
		if h[v] {
			return v
		}
	}
	return -1
}