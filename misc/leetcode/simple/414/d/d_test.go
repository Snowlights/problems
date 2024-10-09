package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxMoves, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-414/problems/maximum-number-of-moves-to-kill-all-pawns/
// https://leetcode.cn/problems/maximum-number-of-moves-to-kill-all-pawns/
