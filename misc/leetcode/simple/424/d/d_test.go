package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minDifference, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-424/problems/minimize-the-maximum-adjacent-element-difference/
// https://leetcode.cn/problems/minimize-the-maximum-adjacent-element-difference/
