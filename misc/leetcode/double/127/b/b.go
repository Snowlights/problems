package main

func minimumLevels(possible []int) int {
	n := len(possible)
	s := 0
	for _, x := range possible {
		s += x
	}
	s = s*2 - n
	pre := 0
	for i, x := range possible[:n-1] {
		pre += x*4 - 2
		if pre > s {
			return i + 1
		}
	}
	return -1
}

func minimumLevels2(possible []int) int {

	for i, v := range possible {
		if v == 0 {
			possible[i] = -1
		} else {
			possible[i] = 1
		}
	}
	pre := make([]int, len(possible)+1)
	for i, v := range possible {
		pre[i+1] = pre[i] + v
	}
	for i := 1; i < len(possible); i++ {
		if pre[len(possible)]-pre[i] < pre[i]-pre[0] {
			return i
		}
	}
	return -1
}
