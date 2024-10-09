package e

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestGetNandResult(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, getNandResult, "e.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
