package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minDamage, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-138/problems/minimum-amount-of-damage-dealt-to-bob/
// https://leetcode.cn/problems/minimum-amount-of-damage-dealt-to-bob/
