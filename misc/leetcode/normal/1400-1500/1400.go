package _400_1500

// 1417
func reformat(s string) string {

	num, word := make([]byte, 0), make([]byte, 0)
	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = append(num, byte(v))
		} else {
			word = append(word, byte(v))
		}
	}

	if abs(len(num)-len(word)) > 1 {
		return ""
	}

	body := make([]byte, len(num)+len(word))
	i, j, k := 0, 0, 0
	if len(word) >= len(num) {
		for i < len(num) && j < len(word) {
			body[k] = word[i]
			k++
			body[k] = num[j]
			k++
			i++
			j++
		}
		if j < len(word) {
			body[k] = word[j]
		}
	} else {
		for i < len(num) && j < len(word) {
			body[k] = num[i]
			k++
			body[k] = word[j]
			k++
			i++
			j++
		}
		if i < len(num) {
			body[k] = num[i]
		}
	}

	return string(body)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
