//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2010/C2
// https://codeforces.com/problemset/status/2010/problem/C2
func TestCF2010C2(t *testing.T) {
	t.Log("Current test is [CF2010C2]")
	testCases := [][2]string{
		{
			`abrakadabrabrakadabra`,
			`YES
			abrakadabra`,
		},
		{
			`acacacaca`,
			`YES
			acacaca`,
		},
		{
			`abcabc`,
			`NO`,
		},
		{
			`abababab`,
			`YES
			ababab`,
		},
		{
			`tatbt`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2010C2, testCases, 0)
}
