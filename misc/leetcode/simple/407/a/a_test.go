// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minChanges, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-407/problems/number-of-bit-changes-to-make-two-integers-equal/
// https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/
