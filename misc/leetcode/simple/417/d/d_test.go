package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, kthCharacter, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-417/problems/find-the-k-th-character-in-string-game-ii/
// https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/
