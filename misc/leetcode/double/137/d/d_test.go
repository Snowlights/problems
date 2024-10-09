package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumValueSum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-137/problems/maximum-value-sum-by-placing-three-rooks-ii/
// https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/
