package main

func perfectPerformance(s string) bool {
	x, y := 0, 0
	for _, v := range s {
		switch v {
		case 'L':
			x--
		case 'R':
			x++
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
