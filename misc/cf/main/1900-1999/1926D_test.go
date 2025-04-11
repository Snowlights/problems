//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1926/D
// https://codeforces.com/problemset/status/1926/problem/D
func TestCF1926D(t *testing.T) {
	t.Log("Current test is [CF1926D]")
	testCases := [][2]string{
		{
			`9
			4
			1 4 3 4
			2
			0 2147483647
			5
			476319172 261956880 2136179468 1671164475 1885526767
			3
			1335890506 811593141 1128223362
			4
			688873446 627404104 1520079543 1458610201
			4
			61545621 2085938026 1269342732 1430258575
			4
			0 0 2147483647 2147483647
			3
			0 0 2147483647
			8
			1858058912 289424735 1858058912 2024818580 1858058912 289424735 122665067 289424735`,
			`4
			1
			3
			2
			2
			3
			2
			2
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1926D, testCases, 0)
}
