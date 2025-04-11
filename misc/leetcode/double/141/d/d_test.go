package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, numberOfWays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-141/problems/find-the-number-of-possible-ways-for-an-event/
// https://leetcode.cn/problems/find-the-number-of-possible-ways-for-an-event/
