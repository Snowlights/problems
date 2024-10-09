package double_93

import (
	"math"
	"sort"
	"strconv"
)

// 1
func maximumValue(strs []string) int {
	ans := 0
	for _, v := range strs {
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			ans = max(ans, len(v))
		} else {
			ans = max(ans, int(vv))
		}
	}
	return ans
}

// 2
func maxStarSum(vals []int, edges [][]int, k int) int {
	g := make([][]int, len(vals))
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		g[from] = append(g[from], vals[to])
		g[to] = append(g[to], vals[from])
	}
	ans := math.MinInt32
	for from, to := range g {
		sort.Ints(to)
		tp := k
		sum := vals[from]
		for i := len(to) - 1; i >= 0 && tp > 0; i-- {
			if to[i] < 0 {
				break
			}
			sum += to[i]
			tp--
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 3
func maxJump(stones []int) int {
	ans := stones[1] - stones[0]
	for i := 2; i < len(stones); i++ {
		ans = max(ans, stones[i]-stones[i-2])
	}
	return ans
}

// 4
func minimumTotalCost(nums1, nums2 []int) (ans int64) {
	var swapCnt, modeCnt, mode int
	cnt := make([]int, len(nums1)+1)
	for i, x := range nums1 {
		if x == nums2[i] {
			ans += int64(i)
			swapCnt++
			cnt[x]++
			if cnt[x] > modeCnt {
				modeCnt, mode = cnt[x], x
			}
		}
	}

	for i, x := range nums1 {
		if modeCnt*2 <= swapCnt {
			break
		}
		if x != nums2[i] && x != mode && nums2[i] != mode {
			ans += int64(i)
			swapCnt++
		}
	}
	if modeCnt*2 > swapCnt {
		return -1
	}
	return
}
