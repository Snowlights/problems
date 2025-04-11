// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc282/submit?taskScreenName=abc282_e
// 对拍：https://atcoder.jp/contests/abc282/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc282_e&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`4 10
			4 2 3 2`,
			`20`,
		},
		{
			`20 100
			29 31 68 20 83 66 23 84 69 96 41 61 83 37 52 71 18 55 40 8`,
			`1733`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/abc282/tasks/abc282_e
