package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxRectangleArea, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-427/problems/maximum-area-rectangle-with-point-constraints-ii/
// https://leetcode.cn/problems/maximum-area-rectangle-with-point-constraints-ii/
