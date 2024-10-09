//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1800/E2
// https://codeforces.com/problemset/status/1800/problem/E2
func TestCF1800E2(t *testing.T) {
	t.Log("Current test is [CF1800E2]")
	testCases := [][2]string{
		{
			`7
			6 3
			talant
			atltna
			7 1
			abacaba
			aaaabbc
			12 6
			abracadabraa
			avadakedavra
			5 3
			accio
			cicao
			5 4
			lumos
			molus
			4 3
			uwjt
			twju
			4 3
			kvpx
			vxpk`,
			`YES
			YES
			NO
			YES
			NO
			YES
			NO
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1800E2, testCases, 0)
}
