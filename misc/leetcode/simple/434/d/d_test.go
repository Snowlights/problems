package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, supersequences, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-434/problems/frequencies-of-shortest-supersequences/
// https://leetcode.cn/problems/frequencies-of-shortest-supersequences/
