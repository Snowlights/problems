package main

import "strings"

func hasMatch(s, p string) bool {
	star := strings.IndexByte(p, '*')
	i := strings.Index(s, p[:star])
	return i >= 0 && strings.Contains(s[i+star:], p[star+1:])
}
