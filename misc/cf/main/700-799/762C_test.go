//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/762/C
// https://codeforces.com/problemset/status/762/problem/C
func TestCF762C(t *testing.T) {
	t.Log("Current test is [CF762C]")
	testCases := [][2]string{
		{
			`hi
			bob`,
			`-`,
		},
		{
			`abca
			accepted`,
			`ac`,
		},
		{
			`abacaba
			abcdcba`,
			`abcba`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF762C, testCases, 0)
}
