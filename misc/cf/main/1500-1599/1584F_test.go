//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1584/F
// https://codeforces.com/problemset/status/1584/problem/F
func TestCF1584F(t *testing.T) {
	t.Log("Current test is [CF1584F]")
	testCases := [][2]string{
		{
			`4
			2
			ABC
			CBA
			2
			bacab
			defed
			3
			abcde
			aBcDe
			ace
			2
			codeforces
			technocup`,
			`1
			A
			0
			
			3
			ace
			3
			coc
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1584F, testCases, 0)
}
