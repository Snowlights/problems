package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumValueSum, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-393/problems/minimum-sum-of-values-by-dividing-array/
// https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array/
