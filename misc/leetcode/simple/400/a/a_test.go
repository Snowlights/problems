// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumChairs, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-400/problems/minimum-number-of-chairs-in-a-waiting-room/
// https://leetcode.cn/problems/minimum-number-of-chairs-in-a-waiting-room/
