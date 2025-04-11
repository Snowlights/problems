package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countComponents, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-145/problems/count-connected-components-in-lcm-graph/
// https://leetcode.cn/problems/count-connected-components-in-lcm-graph/
