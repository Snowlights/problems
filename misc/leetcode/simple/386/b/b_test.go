// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, largestSquareArea, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-386/problems/find-the-largest-area-of-square-inside-two-rectangles/
// https://leetcode.cn/problems/find-the-largest-area-of-square-inside-two-rectangles/
