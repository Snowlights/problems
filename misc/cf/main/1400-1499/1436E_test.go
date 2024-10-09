//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1436/E
// https://codeforces.com/problemset/status/1436/problem/E
func TestCF1436E(t *testing.T) {
	t.Log("Current test is [CF1436E]")
	testCases := [][2]string{
		{
			`3
			1 3 2
			`,
			`3`,
		},
		{
			`5
			1 4 3 1 2
			`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1436E, testCases, 0)
}
