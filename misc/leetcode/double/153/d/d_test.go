package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxActiveSectionsAfterTrade, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-153/problems/maximize-active-section-with-trade-ii/
// https://leetcode.cn/problems/maximize-active-section-with-trade-ii/
