package _200_1300

import "sort"

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

// 1283
func smallestDivisor(nums []int, threshold int) int {
	return sort.Search(1e9, func(i int) bool {
		if i == 0 {
			return false
		}
		val := 0
		for _, v := range nums {
			val += v / i
			if v%i > 0 {
				val++
			}
		}
		return val <= threshold
	})
}

// 1292
func maxSideLength(mat [][]int, threshold int) int {

	m, n := len(mat), len(mat[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i, v := range mat {
		for j, vv := range v {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + vv
		}
	}

	maxSideLengthHelper := func(mid int) bool {
		for i := 1; i+mid-1 <= m; i++ {
			for j := 1; j+mid-1 <= n; j++ {
				diff := pre[i+mid-1][j+mid-1] - pre[i+mid-1][j-1] - pre[i-1][j+mid-1] + pre[i-1][j-1]
				if diff <= threshold {
					return true
				}
			}
		}
		return false
	}

	maxlength := min(m, n)
	l, r := 0, maxlength
	for l < r {
		mid := (l + r + 1) / 2
		if maxSideLengthHelper(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
