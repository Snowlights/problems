package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumSubarrayXor, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-413/problems/maximum-xor-score-subarray-queries/
// https://leetcode.cn/problems/maximum-xor-score-subarray-queries/
