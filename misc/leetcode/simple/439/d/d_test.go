package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, generateString, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-439/problems/lexicographically-smallest-generated-string/
// https://leetcode.cn/problems/lexicographically-smallest-generated-string/
