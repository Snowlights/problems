// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc098/submit?taskScreenName=arc098_a
// 对拍：https://atcoder.jp/contests/arc098/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc098_a&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`5
			WEEWW`,
			`1`,
		},
		{
			`12
			WEWEWEEEWWWE`,
			`4`,
		},
		{
			`8
			WWWWWEEE`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc098/tasks/arc098_a
