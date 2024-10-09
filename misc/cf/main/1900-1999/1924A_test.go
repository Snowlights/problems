//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1924/A
// https://codeforces.com/problemset/status/1924/problem/A
func TestCF1924A(t *testing.T) {
	t.Log("Current test is [CF1924A]")
	testCases := [][2]string{
		{
			`3
			2 2 4
			abba
			2 2 3
			abb
			3 3 10
			aabbccabab`,
			`YES
			NO
			aa
			NO
			ccc`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1924A, testCases, 0)
}
