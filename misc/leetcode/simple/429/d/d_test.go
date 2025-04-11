package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minLength, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-429/problems/smallest-substring-with-identical-characters-ii/
// https://leetcode.cn/problems/smallest-substring-with-identical-characters-ii/
