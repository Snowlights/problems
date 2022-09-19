package _400_1500

import (
	"math/bits"
	"sort"
)

// 1480

func runningSum(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i] + nums[i]
	}
	return pre
}

// 1482
func minDays(bloomDay []int, m, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}
	return sort.Search(maxDay, func(days int) bool {
		flowers, bouquets := 0, 0
		for _, d := range bloomDay {
			if d > days {
				flowers = 0
			} else {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}
		}
		return bouquets >= m
	})
}

// 1484
// SELECT sell_date,
//       COUNT(DISTINCT product) num_sold,
//       GROUP_CONCAT(DISTINCT product) products
// FROM Activities
// GROUP BY sell_date

// 1494
func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	// 计算每门课的先修课集合
	pre := make([]int, n)
	for _, d := range dependencies {
		pre[d[1]-1] |= 1 << (d[0] - 1)
	}
	m := 1 << n
	// 计算所有课程集合的先修课集合，不合法的标记为 -1
	totPre := make([]int, m)
	for i := range totPre {
		if bits.OnesCount(uint(i)) > k {
			totPre[i] = -1
			continue
		}
		for s := uint(i); s > 0; s &= s - 1 {
			p := pre[bits.TrailingZeros(s)]
			if p&i > 0 {
				totPre[i] = -1
				break
			}
			totPre[i] |= p
		}
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for s, v := range dp {
		t := m - 1 ^ s                               // 补集
		for sub := t; sub > 0; sub = (sub - 1) & t { // 枚举下个学期要学的课
			if p := totPre[sub]; p >= 0 && s&p == p { // 这些课的先修课必须合法且在之前的学期里必须上过
				dp[s|sub] = min(dp[s|sub], v+1)
			}
		}
	}
	return dp[m-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
