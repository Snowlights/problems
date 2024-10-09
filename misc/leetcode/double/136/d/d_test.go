package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, timeTaken, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-136/problems/time-taken-to-mark-all-nodes/
// https://leetcode.cn/problems/time-taken-to-mark-all-nodes/
