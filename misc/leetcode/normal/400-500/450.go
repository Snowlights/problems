package _00_500

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 451
func frequencySort(s string) string {
	type pair struct {
		i, count int
	}
	pairList, pairMap := make([]*pair, 0), map[int]*pair{}
	for _, v := range s {
		val := int(v - 'a')
		m, ok := pairMap[val]
		if !ok {
			m = &pair{
				i:     val,
				count: 0,
			}
			pairList = append(pairList, m)
			pairMap[val] = m
		}
		m.count++
	}
	sort.Slice(pairList, func(i int, j int) bool {
		return pairList[i].count > pairList[j].count
	})
	ans := ""
	for _, pair := range pairList {
		ans += strings.Repeat(fmt.Sprintf("%s", string(rune(pair.i+'a'))), pair.count)
	}
	return ans
}

// 456
func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}

// 459
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
