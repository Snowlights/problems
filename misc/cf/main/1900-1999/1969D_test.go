//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1969/D
// https://codeforces.com/problemset/status/1969/problem/D
func TestCF1969D(t *testing.T) {
	t.Log("Current test is [CF1969D]")
	testCases := [][2]string{{
		`4
		2 0
		2 1
		1 2
		4 1
		1 2 1 4
		3 3 2 3
		4 2
		2 1 1 1
		4 2 3 2
		6 2
		1 3 4 9 1 3
		7 6 8 10 6 8`,
		`1
		1
		0
		7`,
	}}
	codeforces.AssertEqualStringCase(t, CF1969D, testCases, 0)
}
