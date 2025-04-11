package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, lengthAfterTransformations, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-421/problems/total-characters-in-string-after-transformations-ii/
// https://leetcode.cn/problems/total-characters-in-string-after-transformations-ii/
