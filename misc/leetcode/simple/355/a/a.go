package main

import "strings"

func splitWordsBySeparator(a []string, separator byte) (ans []string) {
	for _, v := range a {
		for _, vv := range strings.Split(v, string(separator)) {
			if len(vv) > 0 {
				ans = append(ans, vv)
			}
		}
	}
	return
}
