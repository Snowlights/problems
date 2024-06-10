//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1416/D
// https://codeforces.com/problemset/status/1416/problem/D
func TestCF1416D(t *testing.T) {
	t.Log("Current test is [CF1416D]")
	testCases := [][2]string{
		{
			`5 4 6
			1 2 5 4 3
			1 2
			2 3
			1 3
			4 5
			1 1
			2 1
			2 3
			1 1
			1 2
			1 2
			`,
			`5
			1
			2
			0
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1416D, testCases, 0)
}
