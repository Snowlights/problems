package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxCollectedFruits, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-144/problems/find-the-maximum-number-of-fruits-collected/
// https://leetcode.cn/problems/find-the-maximum-number-of-fruits-collected/
