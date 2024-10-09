//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/486/E
// https://codeforces.com/problemset/status/486/problem/E
func TestCF486E(t *testing.T) {
	t.Log("Current test is [CF486E]")
	testCases := [][2]string{
		{
			`1
			4`,
			`3`,
		},
		{
			`4
			1 3 2 5`,
			`3223`,
		},
		{
			`4
			1 5 2 3`,
			`3133`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF486E, testCases, 0)
}
