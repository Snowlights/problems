package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, makeStringGood, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-428/problems/minimum-operations-to-make-character-frequencies-equal/
// https://leetcode.cn/problems/minimum-operations-to-make-character-frequencies-equal/
