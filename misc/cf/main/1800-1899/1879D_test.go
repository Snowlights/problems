//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1879/D
// https://codeforces.com/problemset/status/1879/problem/D
func TestCF1879D(t *testing.T) {
	t.Log("Current test is [CF1879D]")
	testCases := [][2]string{
		{
			`3
			1 3 2`,
			`12`,
		},
		{
			`4
			39 68 31 80`,
			`1337`,
		},
		{
			`7
			313539461 779847196 221612534 488613315 633203958 394620685 761188160`,
			`257421502`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1879D, testCases, 0)
}
