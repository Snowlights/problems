package _100_1200

import (
	"math"
	"sort"
)

// 1162
func maxDistance(grid [][]int) int {
	ans := 0
	a, b := [][]int{}, [][]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				a = append(a, []int{i, j})
			} else {
				b = append(b, []int{i, j})
			}
		}
	}

	if len(a) == 0 || len(b) == 0 {
		return -1
	}

	for _, v := range b {
		tmp := math.MaxInt32
		for _, u := range a {
			val := abs(v[0]-u[0]) + abs(v[1]-u[1])
			tmp = min(tmp, val)
		}
		ans = max(ans, tmp)
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1170
func numSmallerByFrequency(queries []string, words []string) []int {
	w := make([]int, 0, len(words))
	f := func(word string) int {
		h := make(map[int]int)
		for i := range word {
			h[int(word[i]-'a')]++
		}
		for i := 0; i < 26; i++ {
			if val, ok := h[i]; ok {
				return val
			}
		}
		return 0
	}
	for _, word := range words {
		w = append(w, f(word))
	}
	sort.Ints(w)
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = len(q) - sort.SearchInts(w, f(q)+1)
	}
	return ans
}
