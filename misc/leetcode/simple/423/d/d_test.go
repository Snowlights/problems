package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countKReducibleNumbers, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-423/problems/count-k-reducible-numbers-less-than-n/
// https://leetcode.cn/problems/count-k-reducible-numbers-less-than-n/
