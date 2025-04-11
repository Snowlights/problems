//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/903/F
// https://codeforces.com/problemset/status/903/problem/F
func TestCF903F(t *testing.T) {
	t.Log("Current test is [CF903F]")
	testCases := [][2]string{
		{
			`4
			1 10 8 20
			***.
			***.
			***.
			...*`,
			`9`,
		},
		{
			`7
			2 1 8 2
			.***...
			.***..*
			.***...
			....*..`,
			`3`,
		},
		{
			`4
			10 10 1 10
			***.
			*..*
			*..*
			.***`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF903F, testCases, 0)
}
