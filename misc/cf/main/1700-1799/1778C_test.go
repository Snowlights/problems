//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1778/C
// https://codeforces.com/problemset/status/1778/problem/C
func TestCF1778C(t *testing.T) {
	t.Log("Current test is [CF1778C]")
	testCases := [][2]string{
		{
			`6
			3 1
			abc
			abd
			3 0
			abc
			abd
			3 1
			xbb
			xcd
			4 1
			abcd
			axcb
			3 10
			abc
			abd
			10 3
			lkwhbahuqa
			qoiujoncjb`,
			`6
			3
			6
			6
			6
			11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1778C, testCases, 0)
}
