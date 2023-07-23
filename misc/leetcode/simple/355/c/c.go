package main

import "sort"

func maxIncreasingGroups(usageLimits []int) (ans int) {
	sort.Ints(usageLimits)
	s, sum := make([]int, len(usageLimits)), usageLimits[0]
	s[0] = 1
	for i := 1; i < len(usageLimits); i++ {
		sum += usageLimits[i]
		if usageLimits[i] > s[i-1] || (s[i-1]+1)*(s[i-1]+2)/2 <= sum {
			s[i] = s[i-1] + 1
		} else {
			s[i] = s[i-1]
		}
	}

	return s[len(s)-1]
}

// class Solution:
//    def maxIncreasingGroups(self, usageLimits: List[int]) -> int:
//        usageLimits2 = sorted(usageLimits, reverse=False)
//        n = len(usageLimits2)
//        # 满足数字约束
//        x = [0]*n
//        x[0] = 1
//        vsum = usageLimits2[0]
//        for i in range(1,n):
//            vsum += usageLimits2[i]
//            if usageLimits2[i] > x[i-1] or (x[i-1]+1)*(x[i-1]+2)/2 <= vsum:
//                x[i] = x[i-1] + 1
//            else:
//                x[i] = x[i-1]
//
//        return x[-1]
//
