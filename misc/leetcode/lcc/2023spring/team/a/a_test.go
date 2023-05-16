package a

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestRuneReserve(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, runeReserve, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
