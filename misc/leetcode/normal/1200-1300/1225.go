package _200_1300

import "strings"

// 1249
func minRemoveToMakeValid(s string) string {

	b, length := []byte(s), len(s)
	sum := 0
	for i := 0; i < length; i++ {
		if b[i] == '(' {
			sum++
		} else if b[i] == ')' {
			sum--
			if sum < 0 {
				sum = 0
				b[i] = '0'
			}
		}
	}
	sum = 0
	for i := length - 1; i >= 0; i-- {
		if b[i] == ')' {
			sum++
		} else if b[i] == '(' {
			sum--
			if sum < 0 {
				sum = 0
				b[i] = '0'
			}
		}
	}

	return strings.ReplaceAll(string(b), "0", "")
}
