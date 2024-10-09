package simple_329

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func alternateDigitSum(n int) int {
	ans, flag := 0, true
	num := strconv.Itoa(n)
	for _, v := range num {
		if flag {
			ans += int(v - '0')
		} else {
			ans -= int(v - '0')
		}
		flag = !flag
	}
	return ans
}

// 2
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool { return score[i][k] > score[j][k] })
	return score
}

// 3
func makeStringsEqual(s string, target string) bool {
	return strings.Contains(s, "1") == strings.Contains(target, "1")
}

// 4
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)

	for i := range nums {
		cnt, mm := 0, math.MaxInt32
		h := make(map[int]int)
		for j := i; j >= 0; j-- {
			h[nums[j]]++
			if h[nums[j]] == 2 {
				cnt += 2
			} else if h[nums[j]] > 2 {
				cnt++
			}
			mm = min(mm, f[j]+cnt)
		}

		f[i+1] = k + mm
	}

	return f[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
