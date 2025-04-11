package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, distanceSum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-148/problems/manhattan-distances-of-all-arrangements-of-pieces/
// https://leetcode.cn/problems/manhattan-distances-of-all-arrangements-of-pieces/
