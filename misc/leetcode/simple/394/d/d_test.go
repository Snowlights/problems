package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, findAnswer, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-394/problems/find-edges-in-shortest-paths/
// https://leetcode.cn/problems/find-edges-in-shortest-paths/
