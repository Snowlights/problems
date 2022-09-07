package double_86

import "math/bits"

// 1
func findSubarrays(nums []int) bool {
	h := make(map[int]bool)
	for i := 1; i < len(nums); i++ {
		if h[nums[i]+nums[i-1]] {
			return true
		}
		h[nums[i]+nums[i-1]] = true
	}
	return false
}

// 2
func isStrictlyPalindromic(n int) bool {
	return false
}

// 3
func maximumRows(mat [][]int, cols int) (ans int) {
	m, n := len(mat), len(mat[0])
	mask := make([]int, m)
	for i, row := range mat {
		for j, v := range row {
			mask[i] |= v << j
		}
	}
	for set := 0; set < 1<<n; set++ {
		if bits.OnesCount(uint(set)) != cols { // 跳过大小不等于 cols 的集合
			continue
		}
		cnt := 0
		for _, row := range mask {
			if row&set == row { // row 是 set 的子集，所有 1 都被覆盖
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//4
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {

	left, ans, sum, q := 0, 0, 0, []int{}
	for right, v := range chargeTimes {
		for len(q) > 0 && chargeTimes[q[len(q)-1]] <= v {
			q = q[:len(q)-1]
		}

		sum += runningCosts[right]
		q = append(q, right)

		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*int64(sum) > budget {
			if q[0] == left {
				q = q[1:]
			}
			sum -= runningCosts[left]
			left++
		}
		ans = max(ans, right-left+1)
	}

	return ans
}
