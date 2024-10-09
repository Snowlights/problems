//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1336/C
// https://codeforces.com/problemset/status/1336/problem/C
func TestCF1336C(t *testing.T) {
	t.Log("Current test is [CF1336C]")
	testCases := [][2]string{
		{
			`abab
			ba`,
			`12`,
		},
		{
			`defineintlonglong
			signedmain`,
			`0`,
		},
		{
			`rotator
			rotator`,
			`4`,
		},
		{
			`cacdcdbbbb
			bdcaccdbbb`,
			`24`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1336C, testCases, 0)
}
