//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1837/D
// https://codeforces.com/problemset/status/1837/problem/D
func TestCF1837D(t *testing.T) {
	t.Log("Current test is [CF1837D]")
	testCases := [][2]string{
		{
			`4
			8
			((())))(
			4
			(())
			4
			))((
			3
			(()`,
			`2
			2 2 2 1 2 2 2 1
			1
			1 1 1 1
			1
			1 1 1 1
			-1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1837D, testCases, 0)
}
