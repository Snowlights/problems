package main

import (
	"fmt"
	"strings"
)

func mathProblem(cal string) string {

	handle := func(s string) (int, int) {
		i, x, sum, flag := 0, 0, 0, 1

		for i < len(s) {
			switch s[i] {
			case 'x':
				x += flag
				i++
			case '+':
				flag = 1
				i++
			case '-':
				flag = -1
				i++
			default:
				v := 0
				for i < len(s) && s[i] >= '0' && s[i] <= '9' {
					v = v*10 + int(s[i]-'0')
					i++
				}
				if i < len(s) && s[i] == 'x' {
					x += flag * v
					i++
				} else {
					sum += flag * v
				}
			}
		}
		return x, sum
	}
	parts := strings.Split(cal, "=")
	ax, av := handle(parts[0])
	bx, bv := handle(parts[1])

	if ax == bx {
		if av-bv != 0 {
			return "No solution"
		}
		return "Infinite solutions"
	}
	if (bv-av)%(ax-bx) != 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", (bv-av)/(ax-bx))
}
