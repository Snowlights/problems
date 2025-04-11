//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1362/B
// https://codeforces.com/problemset/status/1362/problem/B
func TestCF1362B(t *testing.T) {
	t.Log("Current test is [CF1362B]")
	testCases := [][2]string{{
		`6
		4
		1 0 2 3
		6
		10 7 14 8 3 12
		2
		0 2
		3
		1 2 3
		6
		1 4 6 10 11 12
		2
		0 1023`,
		`1
		4
		2
		-1
		-1
		1023`,
	}}
	codeforces.AssertEqualStringCase(t, CF1362B, testCases, 0)
}
