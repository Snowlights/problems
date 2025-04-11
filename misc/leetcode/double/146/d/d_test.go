package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, subsequencesWithMiddleMode, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-146/problems/subsequences-with-a-unique-middle-mode-i/
// https://leetcode.cn/problems/subsequences-with-a-unique-middle-mode-i/
