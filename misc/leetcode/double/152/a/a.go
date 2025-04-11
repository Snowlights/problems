package main

func totalNumbers(digits []int) int {
	set := map[int]struct{}{}
	for i, a := range digits { // 个位数
		if a%2 > 0 {
			continue
		}
		for j, b := range digits { // 十位数
			if j == i {
				continue
			}
			for k, c := range digits { // 百位数
				if c == 0 || k == i || k == j {
					continue
				}
				set[c*100+b*10+a] = struct{}{}
			}
		}
	}
	return len(set)
}
