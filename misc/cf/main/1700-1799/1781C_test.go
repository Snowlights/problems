//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1781/C
// https://codeforces.com/problemset/status/1781/problem/C
func TestCF1781C(t *testing.T) {
	t.Log("Current test is [CF1781C]")
	testCases := [][2]string{
		{
			`4
			5
			hello
			10
			codeforces
			5
			eevee
			6
			appall`,
			`1
			helno
			2
			codefofced
			1
			eeeee
			0
			appall
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1781C, testCases, 0)
}
