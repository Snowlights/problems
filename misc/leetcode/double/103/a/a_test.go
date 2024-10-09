package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximizeSum, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-103/problems/maximum-sum-with-exactly-k-elements/
// https://leetcode.cn/problems/maximum-sum-with-exactly-k-elements/
