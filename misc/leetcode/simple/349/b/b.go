package main

func smallestString(s string) (ans string) {
	body, i := []byte(s), 0
	for i < len(s) && body[i] == 'a' {
		i++
	}
	flag := false
	for i < len(s) && body[i] != 'a' {
		body[i]--
		flag = true
		i++
	}

	if !flag {
		body[len(body)-1]--
		if body[len(body)-1] < 'a' {
			body[len(body)-1] = 'z'
		}
	}

	return string(body)
}
