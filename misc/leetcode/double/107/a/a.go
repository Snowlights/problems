package main

func maximumNumberOfStringPairs(a []string) (ans int) {

	for i, v := range a {
		for _, vv := range a[i+1:] {
			if v[0] == vv[1] && vv[0] == v[1] {
				ans++
			}
		}
	}

	return
}
