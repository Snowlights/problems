package main

func minimumChairs(s string) (ans int) {
	cnt := 0
	for _, c := range s {
		if c == 'E' {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
		}
	}
	return
}
