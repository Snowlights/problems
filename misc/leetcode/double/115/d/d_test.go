// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, countSubMultisets, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-115/problems/count-of-sub-multisets-with-bounded-sum/
// https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/
