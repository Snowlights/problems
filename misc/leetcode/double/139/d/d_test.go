package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxPathLength, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-139/problems/length-of-the-longest-increasing-path/
// https://leetcode.cn/problems/length-of-the-longest-increasing-path/
