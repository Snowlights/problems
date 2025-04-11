//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1950/E
// https://codeforces.com/problemset/status/1950/problem/E
func TestCF1950E(t *testing.T) {
	t.Log("Current test is [CF1950E]")
	testCases := [][2]string{
		{
			`5
			4
			abaa
			4
			abba
			13
			slavicgslavic
			8
			hshahaha
			20
			stormflamestornflame`,
			`1
			4
			13
			2
			10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1950E, testCases, 0)
}
