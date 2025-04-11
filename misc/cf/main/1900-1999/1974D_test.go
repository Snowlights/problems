//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1974/D
// https://codeforces.com/problemset/status/1974/problem/D
func TestCF1974D(t *testing.T) {
	t.Log("Current test is [CF1974D]")
	testCases := [][2]string{
		{
			`10
			6
			NENSNE
			3
			WWW
			6
			NESSWS
			2
			SN
			2
			WE
			4
			SSNN
			4
			WESN
			2
			SS
			4
			EWNN
			4
			WEWE`,
			`RRHRRH
			NO
			HRRHRH
			NO
			NO
			RHRH
			RRHH
			RH
			RRRH
			RRHH`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1974D, testCases, 0)
}
