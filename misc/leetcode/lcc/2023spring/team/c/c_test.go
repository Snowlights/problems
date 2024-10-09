package c

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestExtractMantra(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, extractMantra, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
