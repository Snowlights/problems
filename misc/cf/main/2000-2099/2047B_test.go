//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2047/B
// https://codeforces.com/problemset/status/2047/problem/B
func TestCF2047B(t *testing.T) {
	t.Log("Current test is [CF2047B]")
	testCases := [][2]string{
		{
			`6
			3
			abc
			4
			xyyx
			8
			alphabet
			1
			k
			10
			aabbccddee
			6
			ttbddq`,
			`cbc
			yyyx
			alphaaet
			k
			eabbccddee
			tttddq`,
		},
		{
			`1
			4
			vjha`,
			`vjhv`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2047B, testCases, 0)
}
