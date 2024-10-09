package main

func canBeEqual(x string, y string) (ans bool) {
	cnt := [2][26]int{}
	for i, b := range x {
		cnt[i&1][b-'a']++
	}
	for i, b := range y {
		cnt[i&1][b-'a']--
	}
	for _, c := range cnt {
		for _, c := range c {
			if c != 0 {
				return
			}
		}
	}
	return true
}
