package main

func finalString(s string) string {
	ans := make([]byte, 0, len(s))
	reserve := func() {
		for i := 0; i < len(ans)/2; i++ {
			ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
		}
	}
	for _, v := range s {
		if v == 'i' {
			reserve()
			continue
		}
		ans = append(ans, byte(v))
	}

	return string(ans)
}
