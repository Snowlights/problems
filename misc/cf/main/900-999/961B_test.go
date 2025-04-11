//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/961/B
// https://codeforces.com/problemset/status/961/problem/B
func TestCF961B(t *testing.T) {
	t.Log("Current test is [CF961B]")
	testCases := [][2]string{
		{
			`6 3
			1 3 5 2 5 4
			1 1 0 1 0 0`,
			`16`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF961B, testCases, 0)
}
