package simple_332

import (
	"sort"
	"strconv"
	"strings"
)

// 1
func findTheArrayConcVal(nums []int) int64 {
	ans := int(0)

	for len(nums) > 0 {
		if len(nums) == 1 {
			ans += nums[0]
			break
		}
		val, _ := strconv.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[len(nums)-1]))
		ans += val
		nums = nums[1 : len(nums)-1]
	}

	return int64(ans)
}

// 2
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := int(0)
	for i, v := range nums {
		r := sort.SearchInts(nums[:i], upper-v+1)
		l := sort.SearchInts(nums[:i], lower-v)
		ans += r - l
	}
	return int64(ans)
}

// 3
func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct {
		l, r int
	}
	h := make(map[int]*pair)
	idx := strings.Index(s, "0")
	if idx < len(s) {
		h[0] = &pair{idx, idx}
	}

	for l := range s {
		if s[l] == '0' {
			continue
		}
		val := 0
		for r := l; r < l+31 && r < len(s); r++ {
			val = val<<1 | int(s[r]&1)
			if _, ok := h[val]; !ok {
				h[val] = &pair{
					l: l,
					r: r,
				}
			}
		}
	}
	ans := make([][]int, len(queries))
	notFound := []int{-1, -1} // 避免重复创建
	for i, q := range queries {
		p, ok := h[q[0]^q[1]]
		if !ok {
			ans[i] = notFound
		} else {
			ans[i] = []int{p.l, p.r}
		}
	}
	return ans
}

// 4
func minimumScore(s, t string) int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1
	}

	if suf[0] == 0 {
		return 0
	}
	ans := suf[0]
	for i, j := 0, 0; i < n; i++ {
		if s[i] == t[j] {
			j++
		}
		ans = min(ans, suf[i+1]-j)
	}

	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
