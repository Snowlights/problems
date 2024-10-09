//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/570/E
// https://codeforces.com/problemset/status/570/problem/E
func TestCF570E(t *testing.T) {
	t.Log("Current test is [CF570E]")
	testCases := [][2]string{
		{
			`3 4
			aaab
			baaa
			abba`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF570E, testCases, 0)
}
