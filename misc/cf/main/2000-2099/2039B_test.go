//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2039/B
// https://codeforces.com/problemset/status/2039/problem/B
func TestCF2039B(t *testing.T) {
	t.Log("Current test is [CF2039B]")
	testCases := [][2]string{
		{
			`5
dcabaac
a
youknowwho
codeforces
bangladesh`,
			`abaa
-1
youknowwho
eforce
bang`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF2039B, testCases, 0)
}
