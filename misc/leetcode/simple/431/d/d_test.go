package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumWeight, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-431/problems/maximum-score-of-non-overlapping-intervals/
// https://leetcode.cn/problems/maximum-score-of-non-overlapping-intervals/
