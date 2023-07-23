package main

import "strings"

func splitWordsBySeparator(a []string, separator byte) (ans []string) {
	tmp := []string{}
	for _, v := range a {
		tmp = append(tmp, strings.Split(v, string(separator))...)
	}
	for _, v := range tmp {
		if len(v) == 0 {
			continue
		}
		ans = append(ans, v)
	}
	return
}
