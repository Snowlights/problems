//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1729/problem/G
// https://codeforces.com/problemset/status/1729/problem/G
func TestCF1729G(t *testing.T) {
	t.Log("Current test is [CF1729G]")
	testCases := [][2]string{
		{
			`8
			abababacababa
			aba
			ddddddd
			dddd
			xyzxyz
			xyz
			abc
			abcd
			abacaba
			abaca
			abc
			def
			aaaaaaaa
			a
			aaaaaaaa
			aa`,
			`2 2
			1 4
			2 1
			0 1
			1 1
			0 1
			8 1
			3 6
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1729G, testCases, 0)
}
