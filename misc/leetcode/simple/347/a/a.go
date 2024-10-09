package main

import (
	"strings"
)

func removeTrailingZeros(n string) (ans string) {
	return strings.TrimRight(n, "0")
}
