// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, numberOfSequence, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-374/problems/count-the-number-of-infection-sequences/
// https://leetcode.cn/problems/count-the-number-of-infection-sequences/
