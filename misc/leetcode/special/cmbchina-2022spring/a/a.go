package main

import "strings"

func deleteText(article string, id int) string {
	if article[id] == ' ' {
		return article
	}
	parts := strings.Split(article, " ")
	cnt := strings.Count(article[:id], " ")
	ans := append(parts[:cnt], parts[cnt+1:]...)
	return strings.Join(ans, " ")
}
