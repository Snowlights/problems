package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minStartingIndex, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-140/problems/find-the-occurrence-of-first-almost-equal-substring/
// https://leetcode.cn/problems/find-the-occurrence-of-first-almost-equal-substring/
