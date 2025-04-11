package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxSubarraySum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-147/problems/maximize-subarray-sum-after-removing-all-occurrences-of-one-element/
// https://leetcode.cn/problems/maximize-subarray-sum-after-removing-all-occurrences-of-one-element/
