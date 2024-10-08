// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc298/submit?taskScreenName=abc298_a
// 对拍：https://atcoder.jp/contests/abc298/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc298_a&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`4
			oo--`,
			`Yes`,
		},
		{
			`3
			---`,
			`No`,
		},
		{
			`1
			o`,
			`Yes`,
		},
		{
			`100
			ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooox`,
			`No`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc298/tasks/abc298_a
