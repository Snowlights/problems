package main

import (
	"sort"
	"strconv"
)

func splitNum(n int) (ans int) {
	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	nums := [2]int{}
	for i, v := range s {
		nums[i%2] = nums[i%2]*10 + int(v-'0')
	}
	return nums[0] + nums[1]
}
