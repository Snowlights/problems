//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/276/C
// https://codeforces.com/problemset/status/276/problem/C
func TestCF276C(t *testing.T) {
	t.Log("Current test is [CF276C]")
	testCases := [][2]string{
		{
			`3 3
			5 3 2
			1 2
			2 3
			1 3`,
			`25`,
		},
		{
			`5 3
			5 2 4 1 3
			1 5
			2 3
			2 3`,
			`33`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF276C, testCases, 0)
}
