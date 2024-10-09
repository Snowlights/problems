package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	t.Log("记得初始化所有全局变量")
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, Constructor, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-344/problems/frequency-tracker/
// https://leetcode.cn/problems/frequency-tracker/
