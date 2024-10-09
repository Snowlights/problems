package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumSumSubsequence, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-399/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
// https://leetcode.cn/problems/maximum-sum-of-subsequence-with-non-adjacent-elements/
