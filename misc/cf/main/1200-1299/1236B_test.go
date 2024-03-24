//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1236/B
// https://codeforces.com/problemset/status/1236/problem/B
func TestCF1236B(t *testing.T) {
	t.Log("Current test is [CF1236B]")
	testCases := [][2]string{
		{
			`1 3`,
			`7`,
		},
		{
			`2 2`,
			`9`,
		},
		{
			`1000000000 1000000000`,
			`751201557`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1236B, testCases, 0)
}
