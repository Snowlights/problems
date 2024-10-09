package main

import "strings"

func minLength(s string) (ans int) {

	for {
		if strings.Contains(s, "AB") {
			s = strings.ReplaceAll(s, "AB", "")
		} else if strings.Contains(s, "CD") {
			s = strings.ReplaceAll(s, "CD", "")
		} else {
			return len(s)
		}
	}

}
