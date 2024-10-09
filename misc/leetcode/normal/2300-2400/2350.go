package _300_2400

// 2364

func countBadPairs(nums []int) int64 {

	h := map[int]int{}
	ans := 0
	for i, v := range nums {
		ans += h[v-i]
		h[v-i]++
	}
	tmp := len(nums) * (len(nums) - 1) / 2
	return int64(tmp - ans)
}
