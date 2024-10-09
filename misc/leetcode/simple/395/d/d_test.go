package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, medianOfUniquenessArray, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-395/problems/find-the-median-of-the-uniqueness-array/
// https://leetcode.cn/problems/find-the-median-of-the-uniqueness-array/
