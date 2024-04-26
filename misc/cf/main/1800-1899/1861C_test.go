//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1861/C
// https://codeforces.com/problemset/status/1861/problem/C
func TestCF1861C(t *testing.T) {
	t.Log("Current test is [CF1861C]")
	testCases := [][2]string{
		{
			`7
			++1
			+++1--0
			+0
			0
			++0-+1-+0
			++0+-1+-0
			+1-+0`,
			`YES
			NO
			NO
			NO
			YES
			NO
			NO
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1861C, testCases, 0)
}
