package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxTotalReward, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-401/problems/maximum-total-reward-using-operations-ii/
// https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/
