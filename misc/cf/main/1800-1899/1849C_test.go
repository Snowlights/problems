//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1849/C
// https://codeforces.com/problemset/status/1849/problem/C
func TestCF1849C(t *testing.T) {
	t.Log("Current test is [CF1849C]")
	testCases := [][2]string{
		{
			`3
			6 5
			101100
			1 2
			1 3
			2 4
			5 5
			1 6
			6 4
			100111
			2 2
			1 4
			1 3
			1 2
			1 1
			0
			1 1`,
			`3
			3
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1849C, testCases, 0)
}
