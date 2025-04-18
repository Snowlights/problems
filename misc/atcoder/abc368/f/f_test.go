// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc368/submit?taskScreenName=abc368_f
// 对拍：https://atcoder.jp/contests/abc368/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc368_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`3
			2 3 4`,
			`Anna`,
		},
		{
			`4
			2 3 4 6`,
			`Bruno`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/abc368/tasks/abc368_f
