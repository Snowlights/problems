package main

func numberOfSpecialChars(word string) int {
	ans, cnt := 0, [26]int{}
	for _, v := range word {
		if 'a' <= v && v <= 'z' {
			vi := int(v - 'a')
			switch cnt[vi] {
			case 0:
				cnt[vi] = 1
			case 2:
				cnt[vi] = 3
			}
		} else {
			vi := int(v - 'A')
			switch cnt[vi] {
			case 0:
				cnt[vi] = 3
			case 1:
				cnt[vi] = 2
			}
		}
	}
	for _, v := range cnt {
		if v == 2 {
			ans++
		}
	}
	return ans
}
