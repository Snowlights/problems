package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minMaxSubarraySum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-433/problems/maximum-and-minimum-sums-of-at-most-size-k-subarrays/
// https://leetcode.cn/problems/maximum-and-minimum-sums-of-at-most-size-k-subarrays/
