package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumScore, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-135/problems/maximum-score-from-grid-operations/
// https://leetcode.cn/problems/maximum-score-from-grid-operations/
