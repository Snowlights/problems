package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, lenOfVDiagonal, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-437/problems/length-of-longest-v-shaped-diagonal-segment/
// https://leetcode.cn/problems/length-of-longest-v-shaped-diagonal-segment/
