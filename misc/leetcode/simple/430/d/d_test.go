package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countGoodArrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-430/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/
// https://leetcode.cn/problems/count-the-number-of-arrays-with-k-matching-adjacent-elements/
