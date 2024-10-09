package main

import "strings"

func furthestDistanceFromOrigin(moves string) int {
	return abs(strings.Count(moves, "R")-strings.Count(moves, "L")) + strings.Count(moves, "_")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
