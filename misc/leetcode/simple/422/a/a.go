package main

func isBalanced(num string) bool {
	s := 0
	for i, b := range num {
		s += (i%2*2 - 1) * int(b-'0')
	}
	return s == 0
}
