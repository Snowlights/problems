package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countNonDecreasingSubarrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-432/problems/count-non-decreasing-subarrays-after-k-operations/
// https://leetcode.cn/problems/count-non-decreasing-subarrays-after-k-operations/
