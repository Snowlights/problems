//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/56/D
// https://codeforces.com/problemset/status/56/problem/D
func TestCF56D(t *testing.T) {
	t.Log("Current test is [CF56D]")
	testCases := [][2]string{
		{
			`ABA
			ABBBA`,
			`2
			INSERT 3 B
			INSERT 4 B`,
		},
		{
			`ACCEPTED
			WRONGANSWER`,
			`10
			REPLACE 1 W
			REPLACE 2 R
			REPLACE 3 O
			REPLACE 4 N
			REPLACE 5 G
			REPLACE 6 A
			INSERT 7 N
			INSERT 8 S
			INSERT 9 W
			REPLACE 11 R`,
		},
		{
			`JOYXNKYPF
			GDV`,
			`9
			REPLACE 1 G
			REPLACE 2 D
			REPLACE 3 V
			DELETE 4
			DELETE 4
			DELETE 4
			DELETE 4
			DELETE 4
			DELETE 4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF56D, testCases, 0)
}
