package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, canReachCorner, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-408/problems/check-if-the-rectangle-corner-is-reachable/
// https://leetcode.cn/problems/check-if-the-rectangle-corner-is-reachable/
