// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumChanges, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-368/problems/minimum-changes-to-make-k-semi-palindromes/
// https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes/
