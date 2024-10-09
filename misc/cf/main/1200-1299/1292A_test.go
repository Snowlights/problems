//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1292/A
// https://codeforces.com/problemset/status/1292/problem/A
func TestCF1292A(t *testing.T) {
	t.Log("Current test is [CF1292A]")
	testCases := [][2]string{
		{
			`5 5
			2 3
			1 4
			2 4
			2 3
			1 4`,
			`Yes
			No
			No
			No
			Yes`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1292A, testCases, 0)
}
