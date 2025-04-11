//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1730/D
// https://codeforces.com/problemset/status/1730/problem/D
func TestCF1730D(t *testing.T) {
	t.Log("Current test is [CF1730D]")
	testCases := [][2]string{
		{
			`7
			3
			cbc
			aba
			5
			abcaa
			cbabb
			5
			abcaa
			cbabz
			1
			a
			a
			1
			a
			b
			6
			abadaa
			adaaba
			8
			abcabdaa
			adabcaba`,
			`YES
			YES
			NO
			YES
			NO
			NO
			YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1730D, testCases, 0)
}
