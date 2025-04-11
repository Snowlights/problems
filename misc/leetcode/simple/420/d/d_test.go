package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, findAnswer, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-420/problems/check-if-dfs-strings-are-palindromes/
// https://leetcode.cn/problems/check-if-dfs-strings-are-palindromes/
