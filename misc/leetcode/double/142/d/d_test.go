package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, possibleStringCount, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-142/problems/find-the-original-typed-string-ii/
// https://leetcode.cn/problems/find-the-original-typed-string-ii/
