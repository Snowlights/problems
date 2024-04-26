//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/870/C
// https://codeforces.com/problemset/status/870/problem/C
func TestCF870C(t *testing.T) {
	t.Log("Current test is [CF870C]")
	testCases := [][2]string{
		{
			`1
			12
			`,
			`3`,
		},
		{
			`2
			6
			8
			`,
			`1
			2
			`,
		},
		{
			`3
			1
			2
			3
			`,
			`-1
			-1
			-1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF870C, testCases, 0)
}
